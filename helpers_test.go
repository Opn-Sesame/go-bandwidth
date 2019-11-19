package bandwidth

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	testAccountID     = "123"
	testApplicationID = "1-2-3-4"
)

func expect(t *testing.T, value interface{}, expected interface{}) {
	if !reflect.DeepEqual(value, expected) {
		t.Errorf("Expected %v  - Got %v (%T)", expected, value, value)
	}
}

func expectNil(t *testing.T, value interface{}) {
	if value != nil {
		t.Errorf("Expected nil  - Got %v", value)
	}
}

func shouldFail(t *testing.T, action func() (interface{}, error)) error {
	_, err := action()
	if err == nil {
		t.Fatal("Should fail here")
		return nil
	}
	return err
}

func getAPI(endpoint string) *Client {
	api, _ := New(Opts{testAccountID, "apiToken", "apiSecret", "test", "password", endpoint, endpoint, nil, true})
	return api
}

func createFakeResponse(body string, statusCode int) *http.Response {
	return &http.Response{StatusCode: statusCode,
		Body: nopCloser{bytes.NewReader([]byte(body))}}
}

type RequestHandler struct {
	PathAndQuery string
	Method       string

	EstimatedContent string
	EstimatedHeaders map[string]string

	HeadersToSend    map[string]string
	ContentToSend    string
	StatusCodeToSend int
}

func startMockServer(t *testing.T, handlers []RequestHandler) (*httptest.Server, *Client) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range handlers {
			if handler.Method == "" {
				handler.Method = http.MethodGet
			}
			if handler.StatusCodeToSend == 0 {
				handler.StatusCodeToSend = http.StatusOK
			}
			if handler.Method == r.Method && handler.PathAndQuery == r.URL.String() {
				if handler.EstimatedContent != "" {
					expect(t, readText(t, r.Body), handler.EstimatedContent)
				}
				if handler.EstimatedHeaders != nil {
					for key, value := range handler.EstimatedHeaders {
						expect(t, r.Header.Get(key), value)
					}
				}
				header := w.Header()
				if handler.HeadersToSend != nil {
					for key, value := range handler.HeadersToSend {
						header.Set(key, value)
					}
				}
				if handler.ContentToSend != "" && header.Get("Content-Type") == "" {
					header.Set("Content-Type", "application/json")
				}
				w.WriteHeader(handler.StatusCodeToSend)
				if handler.ContentToSend != "" {
					fmt.Fprintln(w, handler.ContentToSend)
				}
				return
			}
		}
		t.Logf("Unhandled request %s %s, handlers: \n%+v\n", r.Method, r.URL.String(), handlers)
		w.WriteHeader(http.StatusNotFound)
	}))
	api := getAPI(mockServer.URL)
	return mockServer, api
}

func readText(t *testing.T, r io.Reader) string {
	text, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error("Error on reading content")
		return ""
	}
	return string(text)
}
