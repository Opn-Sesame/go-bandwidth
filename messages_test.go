package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetMessages(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/messages",
		Method:       http.MethodGet,
		ContentToSend: `[{
			"id": "{messageId1}",
			"text": "message1"
		}, {
			"id": "{messageId2}",
			"text": "message2"
		}]`}})
	defer server.Close()
	result, err := api.GetMessages()
	if err != nil {
		t.Error("Failed call of GetMessages()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetMessagesFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/messages",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetMessages() })
}

func TestCreateMessage(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/messages",
		Method:           http.MethodPost,
		EstimatedContent: `{"from":"fromNumber","text":"text","to":"toNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/messages/123"}}})
	defer server.Close()
	id, err := api.CreateMessage(map[string]interface{}{"from": "fromNumber", "to": "toNumber", "text": "text"})
	if err != nil {
		t.Error("Failed call of CreateMessage()")
		return
	}
	expect(t, id, "123")
}

func TestCreateMessageFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/messages",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateMessage(map[string]interface{}{"from": "fromNumber", "to": "toNumber", "text": "text"})
	})
}

func TestGetMessage(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/messages/123",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{messageId1}",
			"text": "message1"
		}`}})
	defer server.Close()
	result, err := api.GetMessage("123")
	if err != nil {
		t.Error("Failed call of GetMessage()")
		return
	}
	expect(t, result["id"], "{messageId1}")
}

func TestGetMessageFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/messages/123",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetMessage("123") })
}
