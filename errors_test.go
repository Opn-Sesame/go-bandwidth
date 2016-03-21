package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetErrors(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/errors",
		Method:       http.MethodGet,
		ContentToSend: `[
		{
			"time" : "2012-11-15T01:30:16.208Z",
			"category" : "unavailable",
			"id" : "{userErrorId1}",
			"code" : "no-callback-for-call"
		},
		{
			"time" : "2012-11-15T01:29:24.512Z",
			"category" : "unavailable",
			"id" : "{userErrorId2}",
			"message" : "No application is configured for number +19195556666",
			"code" : "no-application-for-number"
		}]`}})
	defer server.Close()
	result, err := api.GetErrors()
	if err != nil {
		t.Error("Failed call of GetErrors()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetErrorsWithQuery(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/errors?size=2",
		Method:       http.MethodGet,
		ContentToSend: `[
		{
			"time" : "2012-11-15T01:30:16.208Z",
			"category" : "unavailable",
			"id" : "{userErrorId1}",
			"message" : "message",
			"code" : "no-callback-for-call"
		},
		{
			"time" : "2012-11-15T01:29:24.512Z",
			"category" : "unavailable",
			"id" : "{userErrorId2}",
			"message" : "No application is configured for number +19195556666",
			"code" : "no-application-for-number"
		}]`}})
	defer server.Close()
	result, err := api.GetErrors(map[string]string{"size": "2"})
	if err != nil {
		t.Error("Failed call of GetErrors()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetErrorsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/errors",
		Method:       http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func()(interface{}, error){ return api.GetErrors() })
}


func TestGetError(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/errors/123",
		Method:       http.MethodGet,
		ContentToSend: `{
			"time" : "2012-11-15T01:29:24.512Z",
			"category" : "unavailable",
			"id" : "{userErrorId2}",
			"message" : "No application is configured for number +19195556666",
			"code" : "no-application-for-number"
		}`}})
	defer server.Close()
	result, err := api.GetError("123")
	if err != nil {
		t.Error("Failed call of GetError()")
		return
	}
	expect(t, result.ID, "{userErrorId2}")
}

func TestGetErrorFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/errors/123",
		Method:       http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func()(interface{}, error){ return api.GetError("123") })
}

