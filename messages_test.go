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

func TestGetMessagesWithQuery(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/messages?to=123",
		Method:       http.MethodGet,
		ContentToSend: `[{
			"id": "{messageId1}",
			"text": "message1"
		}, {
			"id": "{messageId2}",
			"text": "message2"
		}]`}})
	defer server.Close()
	result, err := api.GetMessages(&GetMessagesQuery{To: "123"})
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
		EstimatedContent: `{"from":"fromNumber","to":"toNumber","text":"text"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/messages/123"}}})
	defer server.Close()
	id, err := api.CreateMessage(&CreateMessageData{From: "fromNumber", To: "toNumber", Text: "text"})
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
		return api.CreateMessage(&CreateMessageData{From: "fromNumber", To: "toNumber", Text: "text"})
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
	expect(t, result.ID, "{messageId1}")
}

func TestGetMessageFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/messages/123",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetMessage("123") })
}
