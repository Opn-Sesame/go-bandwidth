package bandwidth

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	defaultAccountsEndpoint  = "https://dashboard.bandwidth.com"
	accountsPath             = "/api/accounts/"
	defaultMessagingEndpoint = "https://messaging.bandwidth.com"
	messagingPath            = "/api/v2/users/"
)

type endpointRequest int

const (
	messagingRequest endpointRequest = iota
	accountsRequest
)

// RateLimitError is error for 429 http error
type RateLimitError struct {
	Reset time.Time
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("RateLimitError: reset at %v", e.Reset)
}

// Client is main API object
type Client struct {
	AccountID, APIToken, APISecret, UserName, Password string
	AccountsEndpoint, MessagingEndpoint                string
	HTTPClient                                         *http.Client
}

// New creates new instances of api
// It returns Client instance. Use it to make API calls.
func New(accountID, apiToken, apiSecret, userName, password string, accountsEndpoint, messagingEndpoint *string) (*Client, error) {
	if accountID == "" || apiToken == "" || apiSecret == "" || userName == "" || password == "" {
		return nil, errors.New("missing auth data")
	}

	messaging := defaultMessagingEndpoint
	if messagingEndpoint != nil {
		messaging = *messagingEndpoint
	}

	accounts := defaultAccountsEndpoint
	if accountsEndpoint != nil {
		accounts = *accountsEndpoint
	}

	client := &Client{AccountID: accountID, APIToken: apiToken, APISecret: apiSecret, UserName: userName, Password: password,
		AccountsEndpoint:  accounts + accountsPath + accountID,
		MessagingEndpoint: messaging + messagingPath + accountID + "/messages", HTTPClient: http.DefaultClient}
	return client, nil
}

func (c *Client) createRequest(method, path string, requestType endpointRequest) (*http.Request, error) {
	request, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}
	switch requestType {
	case messagingRequest:
		request.SetBasicAuth(c.APIToken, c.APISecret)
		request.Header.Set("Accept", "application/json")
	default:
		request.SetBasicAuth(c.UserName, c.Password)
		request.Header.Set("Accept", "application/xml")
	}
	request.Header.Set("User-Agent", fmt.Sprintf("go-bandwidth/v2"))
	return request, nil
}

func (c *Client) checkJSONResponse(response *http.Response, responseBody interface{}) (interface{}, http.Header, error) {
	defer response.Body.Close()
	body := responseBody
	if body == nil {
		body = map[string]interface{}{}
	}
	rawJSON, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode >= 200 && response.StatusCode < 400 {
		if len(rawJSON) > 0 {
			err = json.Unmarshal([]byte(rawJSON), &body)
			if err != nil {
				return nil, nil, err
			}
		}
		return body, response.Header, nil
	}
	if response.StatusCode == 429 {
		reset, _ := strconv.ParseInt(response.Header.Get("X-RateLimit-Reset"), 10, 64)
		return nil, nil, &RateLimitError{Reset: time.Unix(int64((reset/1000)+1), 0)}
	}
	errorBody := make(map[string]interface{})
	if len(rawJSON) > 0 {
		err = json.Unmarshal([]byte(rawJSON), &errorBody)
		if err != nil {
			return nil, nil, err
		}
	}
	message := errorBody["message"]
	if message == nil {
		message = errorBody["code"]
	}
	if message == nil {
		return nil, nil, fmt.Errorf("Http code %d", response.StatusCode)
	}
	return nil, nil, errors.New(message.(string))
}

func (c *Client) checkXMLResponse(response *http.Response, responseBody interface{}) (interface{}, http.Header, error) {
	defer response.Body.Close()
	body := responseBody
	if body == nil {
		body = map[string]interface{}{}
	}
	rawXML, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode >= 200 && response.StatusCode < 400 {
		if len(rawXML) > 0 {
			err = xml.Unmarshal([]byte(rawXML), &body)
			if err != nil {
				return nil, nil, err
			}
		}
		return body, response.Header, nil
	}
	if response.StatusCode == 429 {
		reset, _ := strconv.ParseInt(response.Header.Get("X-RateLimit-Reset"), 10, 64)
		return nil, nil, &RateLimitError{Reset: time.Unix(int64((reset/1000)+1), 0)}
	}

	// TODO(bashar): Figure out how to deal with errors and what format they come in.
	return nil, nil, fmt.Errorf("Http code %d", response.StatusCode)
}

func (c *Client) makeRequestInternal(method, path string, requestType endpointRequest, data ...interface{}) (interface{}, http.Header, error) {
	request, err := c.createRequest(method, path, requestType)
	var responseBody interface{}
	if err != nil {
		return nil, nil, err
	}
	if len(data) > 0 {
		responseBody = data[0]
	}
	if len(data) > 1 {
		if method == "GET" {
			var item map[string]string
			if data[1] == nil {
				item = make(map[string]string)
			} else {
				var ok bool
				item, ok = data[1].(map[string]string)
				if !ok {
					item = make(map[string]string)
					structType := reflect.TypeOf(data[1]).Elem()
					structValue := reflect.ValueOf(data[1])
					if !structValue.IsNil() {
						structValue = structValue.Elem()
						fieldCount := structType.NumField()
						for i := 0; i < fieldCount; i++ {
							fieldName := structType.Field(i).Name
							fieldValue := structValue.Field(i).Interface()
							if fieldValue == reflect.Zero(structType.Field(i).Type).Interface() {
								//ignore fields with default values
								continue
							}
							item[strings.Replace(strings.ToLower(string(fieldName[0]))+fieldName[1:], "ID", "Id", -1)] = fmt.Sprintf("%v", fieldValue)
						}
					}
				}
			}
			query := make(url.Values)
			for key, value := range item {
				query[key] = []string{value}
			}
			request.URL.RawQuery = query.Encode()
		} else {
			request.Header.Set("Content-Type", "application/json")
			rawJSON, err := json.Marshal(data[1])
			if err != nil {
				return nil, nil, err
			}
			request.Body = nopCloser{bytes.NewReader(rawJSON)}
		}
	}
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, nil, err
	}

	switch requestType {
	case messagingRequest:
		return c.checkJSONResponse(response, responseBody)
	default:
		return c.checkXMLResponse(response, responseBody)
	}
}

func (c *Client) makeMessagingRequest(method, path string, data ...interface{}) (interface{}, http.Header, error) {
	return c.makeRequestInternal(method, path, messagingRequest, data...)
}

func (c *Client) makeAccountsRequest(method, path string, data ...interface{}) (interface{}, http.Header, error) {
	return c.makeRequestInternal(method, path, accountsRequest, data...)
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
