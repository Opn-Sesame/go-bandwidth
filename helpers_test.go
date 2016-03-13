package bandwidth

import (
	"bytes"
	"fmt"
	"reflect"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func expect(t *testing.T, value interface{}, expected interface{}) {
	if !reflect.DeepEqual(value, expected) {
		t.Errorf("Expected %v  - Got %v", expected, value)
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

func getAPI() *Client {
	api, _ := New("userId", "apiToken", "apiSecret")
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
	api := getAPI()
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range handlers {
			if handler.Method == "" {
				handler.Method = http.MethodGet
			}
			if handler.StatusCodeToSend == 0 {
				handler.StatusCodeToSend = http.StatusOK
			}
			if handler.Method == r.Method && handler.PathAndQuery == r.URL.String() {
				w.WriteHeader(handler.StatusCodeToSend)
				if handler.EstimatedContent != "" {
					expect(t, readText(t, r.Body), handler.EstimatedContent)
				}
				if handler.EstimatedHeaders != nil {
					for key, value := range handler.EstimatedHeaders {
						expect(t, r.Header.Get(key), value)
					}
				}
				if handler.HeadersToSend != nil {
					header := w.Header()
					for key, value := range handler.HeadersToSend {
						header.Set(key, value)
					}
					if handler.ContentToSend != "" && header.Get("Content-Type") == "" {
						header.Set("Content-Type", "application/json")
					}
				}
				if handler.ContentToSend != "" {
					fmt.Fprintln(w, handler.ContentToSend)
				}
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	api.APIEndPoint = mockServer.URL
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
