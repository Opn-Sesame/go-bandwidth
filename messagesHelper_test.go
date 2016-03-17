package bandwidth

import (
	"net/http"
	"testing"
)

func TestSendMessageTo(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/messages",
		Method:           http.MethodPost,
		EstimatedContent: `{"from":"fromNumber","text":"text","to":"toNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/messages/123"}}})
	defer server.Close()
	id, err := api.SendMessageTo("fromNumber", "toNumber", "text")
	if err != nil {
		t.Error("Failed message of SendMessageTo()")
		return
	}
	expect(t, id, "123")
}

func TestSendMessageToWithOptions(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/messages",
		Method:           http.MethodPost,
		EstimatedContent: `{"callbackUrl":"url","from":"fromNumber","text":"text","to":"toNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/messages/123"}}})
	defer server.Close()
	id, err := api.SendMessageTo("fromNumber", "toNumber", "text", map[string]interface{}{"callbackUrl": "url"})
	if err != nil {
		t.Error("Failed message of SendMessageTo()")
		return
	}
	expect(t, id, "123")
}
