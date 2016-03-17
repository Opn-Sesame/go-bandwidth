package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetNumberInfo(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/phoneNumbers/numberInfo/123",
		Method:       http.MethodGet,
		ContentToSend: `{
		"created": "2013-09-23T16:31:15Z",
		"name": "Name",
		"number": "123",
		"updated": "2013-09-23T16:42:18Z"
		}`}})
	defer server.Close()
	result, err := api.GetNumberInfo("123")
	if err != nil {
		t.Error("Failed message of GetNumberInfo()")
		return
	}
	expect(t, result["number"], "123")
}

func TestGetNumberInfoFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/phoneNumbers/numberInfo/123",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetNumberInfo("123") })
}
