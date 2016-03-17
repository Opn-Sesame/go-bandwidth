package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetPhoneNumbers(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/phoneNumbers",
		Method:       http.MethodGet,
		ContentToSend: `[{
			"id": "{phoneNumberId1}",
			"number": "phoneNumber1"
		}, {
			"id": "{phoneNumberId2}",
			"number": "phoneNumber2"
		}]`}})
	defer server.Close()
	result, err := api.GetPhoneNumbers()
	if err != nil {
		t.Error("Failed phoneNumber of GetPhoneNumbers()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetPhoneNumbersFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/phoneNumbers",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetPhoneNumbers() })
}

func TestCreatePhoneNumber(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/phoneNumbers",
		Method:           http.MethodPost,
		EstimatedContent: `{"number":"phoneNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/phoneNumbers/123"}}})
	defer server.Close()
	id, err := api.CreatePhoneNumber(map[string]interface{}{"number": "phoneNumber"})
	if err != nil {
		t.Error("Failed phoneNumber of CreatePhoneNumber()")
		return
	}
	expect(t, id, "123")
}

func TestCreatePhoneNumberFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/phoneNumbers",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreatePhoneNumber(map[string]interface{}{"number": "phoneNumber"})
	})
}

func TestGetPhoneNumber(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/phoneNumbers/123",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{phoneNumberId1}",
			"number": "phoneNumber1"
		}`}})
	defer server.Close()
	result, err := api.GetPhoneNumber("123")
	if err != nil {
		t.Error("Failed phoneNumber of GetPhoneNumber()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetPhoneNumberFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/phoneNumbers/123",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetPhoneNumber("123") })
}


func TestUpdatePhoneNumber(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/phoneNumbers/123",
		EstimatedContent: `{"applicationId":"appId"}`,
		Method:       http.MethodPost}})
	defer server.Close()
	err := api.UpdatePhoneNumber("123", map[string]interface{}{"applicationId": "appId"})
	if err != nil {
		t.Error("Failed phoneNumber of UpdatePhoneNumber()")
		return
	}
}


func TestDeletePhoneNumber(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/phoneNumbers/123",
		Method:       http.MethodDelete}})
	defer server.Close()
	err := api.DeletePhoneNumber("123")
	if err != nil {
		t.Error("Failed phoneNumber of DeletePhoneNumber()")
		return
	}
}
