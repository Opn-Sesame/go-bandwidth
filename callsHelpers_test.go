package bandwidth

import (
	"net/http"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	item := map[string]interface{}{"field1": "value1"}
	mergeMaps(item, map[string]interface{}{"field2": "value2", "field3": "value3"})
	expect(t, item, map[string]interface{}{"field1": "value1", "field2": "value2", "field3": "value3"})
	mergeMaps(item, nil)
	expect(t, item, map[string]interface{}{"field1": "value1", "field2": "value2", "field3": "value3"})
}

func TestCallTo(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls",
		Method:           http.MethodPost,
		EstimatedContent: `{"from":"fromNumber","to":"toNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/calls/123"}}})
	defer server.Close()
	id, err := api.CallTo("fromNumber", "toNumber")
	if err != nil {
		t.Error("Failed call of CallTo()")
		return
	}
	expect(t, id, "123")
}

func TestCallToWithOptions(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls",
		Method:           http.MethodPost,
		EstimatedContent: `{"callbackUrl":"url","from":"fromNumber","to":"toNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/calls/123"}}})
	defer server.Close()
	id, err := api.CallTo("fromNumber", "toNumber", map[string]interface{}{"callbackUrl": "url"})
	if err != nil {
		t.Error("Failed call of CallTo()")
		return
	}
	expect(t, id, "123")
}


func TestAnswerIncomingCall(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"active"}`}})
	defer server.Close()
	err := api.AnswerIncomingCall("123")
	if err != nil {
		t.Error("Failed call of AnswerIncomingCall()")
		return
	}
}

func TestRejectIncomingCall(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"rejected"}`}})
	defer server.Close()
	err := api.RejectIncomingCall("123")
	if err != nil {
		t.Error("Failed call of RejectIncomingCall()")
		return
	}
}

func TestHangUpCall(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"completed"}`}})
	defer server.Close()
	err := api.HangUpCall("123")
	if err != nil {
		t.Error("Failed call of HangUpCall()")
		return
	}
}

func TestSetCallRecodingEnabled(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"recordingEnabled":true}`}})
	defer server.Close()
	err := api.SetCallRecodingEnabled("123", true)
	if err != nil {
		t.Error("Failed call of SetCallRecodingEnabled()")
		return
	}
}

func TestTransferCallTo(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"transferring","transferTo":"number"}`}})
	defer server.Close()
	err := api.TransferCallTo("123", "number")
	if err != nil {
		t.Error("Failed call of TransferCallTo()")
		return
	}
}

func TestTransferCallToWithOptions(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"callbackUrl":"url","state":"transferring","transferTo":"number"}`}})
	defer server.Close()
	err := api.TransferCallTo("123", "number", map[string]interface{}{"callbackUrl": "url"})
	if err != nil {
		t.Error("Failed call of TransferCallTo()")
		return
	}
}

func TestStopGather(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/gather/456",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"completed"}`}})
	defer server.Close()
	err := api.StopGather("123", "456")
	if err != nil {
		t.Error("Failed call of StopGather()")
		return
	}
}

func TestSendDTMFCharactersToCalll(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/dtmf",
		Method:           http.MethodPost,
		EstimatedContent: `{"dtmfOut":"1234"}`}})
	defer server.Close()
	err := api.SendDTMFCharactersToCall("123", "1234")
	if err != nil {
		t.Error("Failed call of SendDTMFCharactersToCalll()")
		return
	}
}
