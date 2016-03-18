package bandwidth

import (
	"net/http"
	"testing"
)

func TestTerminateConference(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"completed"}`}})
	defer server.Close()
	err := api.TerminateConference("123")
	if err != nil {
		t.Error("Failed call of TerminateConference()")
	}
}

func TestMuteConference(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"mute":"true"}`}})
	defer server.Close()
	err := api.MuteConference("123", true)
	if err != nil {
		t.Error("Failed call of MuteConference()")
	}
}

func TestDeleteConferenceMember(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members/456",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"completed"}`}})
	defer server.Close()
	err := api.DeleteConferenceMember("123", "456")
	if err != nil {
		t.Error("Failed call of DeleteConferenceMember()")
	}
}

func TestMuteConferenceMember(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members/456",
		Method:           http.MethodPost,
		EstimatedContent: `{"mute":"true"}`}})
	defer server.Close()
	err := api.MuteConferenceMember("123", "456", true)
	if err != nil {
		t.Error("Failed call of MuteConferenceMember()")
	}
}

func TestHoldConferenceMember(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members/456",
		Method:           http.MethodPost,
		EstimatedContent: `{"hold":"true"}`}})
	defer server.Close()
	err := api.HoldConferenceMember("123", "456", true)
	if err != nil {
		t.Error("Failed call of HoldConferenceMember()")
	}
}
