package bandwidth

import (
	"net/http"
	"testing"
)

func TestCreateConference(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences",
		Method:           http.MethodPost,
		EstimatedContent: `{"from":"fromNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/conferences/123"}}})
	defer server.Close()
	id, err := api.CreateConference(&CreateConferenceData{From: "fromNumber"})
	if err != nil {
		t.Error("Failed call of CreateConference()")
		return
	}
	expect(t, id, "123")
}

func TestCreateConferenceFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateConference(&CreateConferenceData{From: "fromNumber"})
	})
}

func TestGetConference(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/conferences/123",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{conferenceId}",
			"state": "completed"
		}`}})
	defer server.Close()
	result, err := api.GetConference("123")
	if err != nil {
		t.Error("Failed call of GetConference()")
		return
	}
	expect(t, result.ID, "{conferenceId}")
}

func TestGetConferenceFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetConference("123") })
}

func TestUpdateConference(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"completed"}`}})
	defer server.Close()
	err := api.UpdateConference("123",&UpdateConferenceData{State: "completed"})
	if err != nil {
		t.Error("Failed call of UpdateConference()")
		return
	}
}

func TestPlayAudioToConference(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/audio",
		Method:           http.MethodPost,
		EstimatedContent: `{"fileUrl":"file.mp3"}`}})
	defer server.Close()
	err := api.PlayAudioToConference("123", &PlayAudioData{FileURL: "file.mp3"})
	if err != nil {
		t.Error("Failed call of PlayAudioToConference()")
		return
	}
}

func TestGetConferenceMembers(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/conferences/123/members",
		Method:       http.MethodGet,
		ContentToSend: `[
		{
			"id": "{memberId1}"
		},
		{
			"id": "{memberId2}"
		}]`}})
	defer server.Close()
	result, err := api.GetConferenceMembers("123")
	if err != nil {
		t.Error("Failed call of GetConferenceMembers()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetConferenceMembersFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetConferenceMembers("123") })
}

func TestGetConferenceMember(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/conferences/123/members/456",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{member1}"
		}`}})
	defer server.Close()
	result, err := api.GetConferenceMember("123", "456")
	if err != nil {
		t.Error("Failed call of GetConferenceMember()")
		return
	}
	expect(t, result.ID, "{member1}")
}

func TestGetConferenceMemberFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members/456",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetConferenceMember("123", "456") })
}

func TestCreateConferenceMember(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members",
		Method:           http.MethodPost,
		EstimatedContent: `{"callId":"callId"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/conferences/123/members/456"}}})
	defer server.Close()
	id, err := api.CreateConferenceMember("123", &CreateConferenceMemberData{CallID: "callId"})
	if err != nil {
		t.Error("Failed call of CreateConferenceMember()")
		return
	}
	expect(t, id, "456")
}

func TestCreateConferenceMemberFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateConferenceMember("123", &CreateConferenceMemberData{CallID: "callId"})
	})
}

func TestUpdateConferenceMember(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members/456",
		Method:           http.MethodPost,
		EstimatedContent: `{"mute":"true"}`}})
	defer server.Close()
	err := api.UpdateConferenceMember("123", "456", &UpdateConferenceMemberData{Mute: true})
	if err != nil {
		t.Error("Failed call of UpdateConferenceMember()")
	}
}

func TestUpdateConferenceMemberFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members/456",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	err := api.UpdateConferenceMember("123", "456", &UpdateConferenceMemberData{Mute: true})
	if err == nil {
		t.Error("Should fail here")
	}
}

func TestPlayAudioToConferenceMember(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/conferences/123/members/456/audio",
		Method:           http.MethodPost,
		EstimatedContent: `{"fileUrl":"file.mp3"}`}})
	defer server.Close()
	err := api.PlayAudioToConferenceMember("123", "456", &PlayAudioData{FileURL: "file.mp3"})
	if err != nil {
		t.Error("Failed call of PlayAudioToConferenceMember()")
		return
	}
}

func TestConferenceMemberGetCallID(t *testing.T) {
	member := &ConferenceMember{Call: "http://host/123"}
	expect(t, member.GetCallID(), "123")
}
