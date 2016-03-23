package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetApplications(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/applications",
		Method:       http.MethodGet,
		ContentToSend: `[
			{
				"id": "{applicationId}",
				"name": "MyFirstApp",
				"incomingCallUrl": "http://example.com/calls.php",
				"incomingMessageUrl": "http://example.com/messages.php",
				"autoAnswer": true
			},
			{
				"id": "{applicationId}",
				"name": "MySecondApp",
				"incomingCallUrl": "http://example.com/app2/calls.php",
				"incomingMessageUrl": "http://example.com/app2/messages.php",
				"autoAnswer": false
			}
		]`}})
	defer server.Close()
	result, err := api.GetApplications()
	if err != nil {
		t.Errorf("Failed call of GetApplications() %s", err.Error())
		return
	}
	expect(t, len(result), 2)
}

func TestGetApplicationsWithQuery(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/applications?size=2",
		Method:       http.MethodGet,
		ContentToSend: `[
			{
				"id": "{applicationId}",
				"name": "MyFirstApp",
				"incomingCallUrl": "http://example.com/calls.php",
				"incomingMessageUrl": "http://example.com/messages.php",
				"autoAnswer": true
			}
		]`}})
	defer server.Close()
	result, err := api.GetApplications(&GetApplicationQuery{Size: 2})
	if err != nil {
		t.Errorf("Failed call of GetApplications() %s", err.Error())
		return
	}
	expect(t, len(result), 1)
}

func TestGetApplicationsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/applications",
		Method:       http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func()(interface{}, error){ return api.GetApplications() })
}

func TestCreateApplication(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/applications",
		Method:       http.MethodPost,
		EstimatedContent: `{"name":"MyFirstApp","incomingCallUrl":"http://example.com/calls.php"}`,
		HeadersToSend: map[string]string{"Location": "/v1/users/{userId}/applications/123"}}})
	defer server.Close()
	id, err := api.CreateApplication(&ApplicationData{
		Name: "MyFirstApp",
		IncomingCallURL: "http://example.com/calls.php"})
	if err != nil {
		t.Error("Failed call of CreateApplication()")
		return
	}
	expect(t, id, "123")
}

func TestCreateApplicationFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/applications",
		Method:       http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func()(interface{}, error){ return api.CreateApplication(&ApplicationData{
		Name: "MyFirstApp",
		IncomingCallURL: "http://example.com/calls.php"}) })
}

func TestGetApplication(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/applications/123",
		Method:       http.MethodGet,
		ContentToSend: `{
				"id": "{applicationId}",
				"name": "MyFirstApp",
				"incomingCallUrl": "http://example.com/calls.php",
				"incomingMessageUrl": "http://example.com/messages.php",
				"autoAnswer": true
		}`}})
	defer server.Close()
	result, err := api.GetApplication("123")
	if err != nil {
		t.Error("Failed call of GetApplication()")
		return
	}
	expect(t, result.ID, "{applicationId}")
}

func TestGetApplicationFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/applications/123",
		Method:       http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func()(interface{}, error){ return api.GetApplication("123") })
}

func TestUpdateApplication(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/applications/123",
		Method:       http.MethodPost,
		EstimatedContent: `{"incomingCallUrl":"http://example.com/calls.php"}`}})
	defer server.Close()
	err := api.UpdateApplication("123", &ApplicationData{IncomingCallURL: "http://example.com/calls.php"})
	if err != nil {
		t.Error("Failed call of UpdateApplication()")
		return
	}
}

func TestDeleteApplication(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/applications/123",
		Method:       http.MethodDelete}})
	defer server.Close()
	err := api.DeleteApplication("123")
	if err != nil {
		t.Error("Failed call of DeleteApplication()")
		return
	}
}

