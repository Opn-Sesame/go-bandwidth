package bandwidth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client is main API object
type Client struct {
	UserID, APIToken, APISecret string
	APIVersion                  string
	APIEndPoint                 string
}

// New creates new instances of api
func New(userID, apiToken, apiSecret string, other ...string) (*Client, error) {
	apiVersion := "v1"
	apiEndPoint := "https://api.catapult.inetwork.com"
	if userID == "" || apiToken == "" || apiSecret == "" {
		return nil, errors.New("Missing auth data. Please use api := bandwidth.New(\"user-id\", \"api-token\", \"api-secret\")")
	}
	l := len(other)
	if l > 1 {
		apiEndPoint = other[1]
	}
	if l > 0 {
		apiVersion = other[0]
	}
	client := &Client{userID, apiToken, apiSecret, apiVersion, apiEndPoint}
	return client, nil
}

func (c *Client) concatUserPath(path string) string {
	if path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("/users/%s%s", c.UserID, path)
}

func (c *Client) prepareURL(path string) string {
	if path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s/%s%s", c.APIEndPoint, c.APIVersion, path)
}

func (c *Client) createRequest(method, path string) (*http.Request, error) {
	request, err := http.NewRequest(method, c.prepareURL(path), nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(c.APIToken, c.APISecret)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", fmt.Sprintf("go-bandwidth-v%s", Version))
	return request, nil
}

func (c *Client) checkResponse(response *http.Response, defaultResponseBody interface{}) (interface{}, http.Header, error) {
	defer response.Body.Close()
	var body interface{}
	rawJSON, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	if len(rawJSON) > 0 {
		err = json.Unmarshal([]byte(rawJSON), &body)
		if err != nil {
			return nil, nil, err
		}
	}
	if response.StatusCode >= 200 && response.StatusCode < 400 {
		if body == nil {
			return defaultResponseBody, response.Header, nil
		}
		result := body
		switch body.(type) {
			case []interface{}:
				source := body.([]interface{})
				l := len(source)
				list := make([]map[string]interface{}, l)
				for i := 0; i < l; i++ {
					list[i] = source[i].(map[string]interface{})
				}
				result = list
		}
		return result, response.Header, nil
	}
	errorBody := make(map[string]interface{})
	if body != nil {
		errorBody = body.(map[string]interface{})
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

func (c *Client) makeRequest(method, path string, data ...interface{}) (interface{}, http.Header, error) {
	request, err := c.createRequest(method, path)
	var defaultResponseBody interface{}
	defaultResponseBody = nil
	if err != nil {
		return nil, nil, err
	}
	if len(data) > 0 {
		if method == "GET" {
			var item map[string]string
			if data[0] == nil {
				item = make(map[string]string)
			} else {
				item = data[0].(map[string]string)
			}
			query := make(url.Values)
			for key, value := range item {
				query[key] = []string{value}
			}
			request.URL.RawQuery = query.Encode()
		} else {
			request.Header.Set("Content-Type", "application/json")
			rawJSON, err := json.Marshal(data[0])
			if err != nil {
				return nil, nil, err
			}
			request.Body = nopCloser{bytes.NewReader(rawJSON)}
		}
	}
	if len(data) > 1 {
		defaultResponseBody = data[1]
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	return c.checkResponse(response, defaultResponseBody)
}

func getIDFromLocationHeader(headers http.Header) string{
	return getIDFromLocation(headers.Get("Location"))
}

func getIDFromLocation(location string) string{
	list := strings.Split(location, "/")
	l := len(list)
	if l > 0 {
		return list[l - 1]
	}
	return ""
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
