package bandwidth

import (
	"bytes"
	"context"
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

// Opts are the options to create the client.
type Opts struct {
	// mandatory options.
	AccountID, APIToken, APISecret, UserName, Password string
	//optional
	AccountsEndpoint, MessagingEndpoint string
	HTTPClient                          *http.Client
}

// Client is main API object
type Client struct {
	accountID, apiToken, apiSecret, userName, password string
	AccountsEndpoint, MessagingEndpoint                string
	httpClient                                         *http.Client
}

// New creates new instances of api
// It returns Client instance. Use it to make API calls.
func New(opts Opts) (*Client, error) {
	if opts.AccountID == "" || opts.APIToken == "" || opts.APISecret == "" ||
		opts.UserName == "" || opts.Password == "" {
		return nil, errors.New("missing auth data")
	}

	messaging := defaultMessagingEndpoint
	if opts.MessagingEndpoint != "" {
		messaging = opts.MessagingEndpoint
	}

	accounts := defaultAccountsEndpoint
	if opts.AccountsEndpoint != "" {
		accounts = opts.AccountsEndpoint
	}

	client := http.DefaultClient
	if opts.HTTPClient != nil {
		client = opts.HTTPClient
	}

	c := &Client{accountID: opts.AccountID, apiToken: opts.APIToken, apiSecret: opts.APISecret,
		userName: opts.UserName, password: opts.Password,
		AccountsEndpoint:  accounts + accountsPath + opts.AccountID,
		MessagingEndpoint: messaging + messagingPath + opts.AccountID + "/messages", httpClient: client}
	return c, nil
}

func (c *Client) createRequest(ctx context.Context, method, path string, requestType endpointRequest) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}
	switch requestType {
	case messagingRequest:
		request.SetBasicAuth(c.apiToken, c.apiSecret)
		request.Header.Set("Accept", "application/json")
	default:
		request.SetBasicAuth(c.userName, c.password)
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

func (c *Client) makeRequestInternal(ctx context.Context, method, path string, requestType endpointRequest, data ...interface{}) (interface{}, http.Header, error) {
	request, err := c.createRequest(ctx, method, path, requestType)
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
			var body []byte
			var err error
			switch requestType {
			case messagingRequest:
				request.Header.Set("Content-Type", "application/json")
				body, err = json.Marshal(data[1])
			default:
				request.Header.Set("Content-Type", "application/xml")
				body, err = xml.Marshal(data[1])
			}
				if err != nil {
					return nil, nil, err
				}
				request.Body = nopCloser{bytes.NewReader(body)}

		}
	}
	response, err := c.httpClient.Do(request)
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

func (c *Client) makeMessagingRequest(ctx context.Context, method, path string, data ...interface{}) (interface{}, http.Header, error) {
	return c.makeRequestInternal(ctx, method, path, messagingRequest, data...)
}

func (c *Client) makeAccountsRequest(ctx context.Context, method, path string, data ...interface{}) (interface{}, http.Header, error) {
	return c.makeRequestInternal(ctx, method, path, accountsRequest, data...)
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
