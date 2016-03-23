package bandwidth

import (
	"net/http"
	"testing"
)


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
		EstimatedContent: `{"recordingEnabled":"true"}`}})
	defer server.Close()
	err := api.SetCallRecodingEnabled("123", true)
	if err != nil {
		t.Error("Failed call of SetCallRecodingEnabled()")
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
