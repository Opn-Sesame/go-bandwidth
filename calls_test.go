package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetCalls(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/calls",
		Method:       http.MethodGet,
		ContentToSend: `[{
			"id": "{callId1}",
			"direction": "out",
			"from": "{fromNumber}",
			"to": "{number}",
			"recordingEnabled": false,
			"callbackUrl": "",
			"state": "completed",
			"startTime": "2013-02-08T13:15:47.587Z",
			"activeTime": "2013-02-08T13:15:52.347Z",
			"endTime": "2013-02-08T13:15:55.887Z",
			"chargeableDuration": 60,
			"events": "https://.../calls/{callId1}/events"
		}, {
			"id": "{callId2}",
			"direction": "out",
			"from": "{number}",
			"to": "{toNumber}",
			"recordingEnabled": false,
			"callbackUrl": "",
			"state": "active",
			"startTime": "2013-02-08T13:15:47.587Z",
			"activeTime": "2013-02-08T13:15:52.347Z",
			"events": "https://.../calls/{callId2}/events"
		}]`}})
	defer server.Close()
	result, err := api.GetCalls()
	if err != nil {
		t.Error("Failed call of GetCalls()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetCallsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetCalls() })
}

func TestCreateCall(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls",
		Method:           http.MethodPost,
		EstimatedContent: `{"from":"fromNumber","to":"toNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/calls/123"}}})
	defer server.Close()
	id, err := api.CreateCall(&CreateCallData{
		From: "fromNumber",
		To:   "toNumber"})
	if err != nil {
		t.Error("Failed call of CreateCall()")
		return
	}
	expect(t, id, "123")
}

func TestCreateCallFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateCall(&CreateCallData{
			From: "fromNumber",
			To:   "toNumber"})
	})
}

func TestGetCall(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/calls/123",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{callId}",
			"state": "completed"
		}`}})
	defer server.Close()
	result, err := api.GetCall("123")
	if err != nil {
		t.Error("Failed call of GetCall()")
		return
	}
	expect(t, result.ID, "{callId}")
}

func TestGetCallFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetCall("123") })
}

func TestUpdateCall(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"completed"}`}})
	defer server.Close()
	_, err := api.UpdateCall("123", &UpdateCallData{State: "completed"})
	if err != nil {
		t.Error("Failed call of UpdateCall()")
		return
	}
}

func TestUpdateCallWithLocationHeader(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123",
		Method:           http.MethodPost,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/calls/456"},
		EstimatedContent: `{"state":"completed"}`}})
	defer server.Close()
	id, err := api.UpdateCall("123", &UpdateCallData{State: "completed"})
	if err != nil {
		t.Error("Failed call of UpdateCall()")
		return
	}
	expect(t, id, "456")
}

func TestPlayAudioToCall(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/audio",
		Method:           http.MethodPost,
		EstimatedContent: `{"fileUrl":"file.mp3"}`}})
	defer server.Close()
	err := api.PlayAudioToCall("123", &PlayAudioData{FileURL: "file.mp3"})
	if err != nil {
		t.Error("Failed call of PlayAudioToCall()")
		return
	}
}

func TestSendDTMFToCall(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/dtmf",
		Method:           http.MethodPost,
		EstimatedContent: `{"dtmfOut":"1234"}`}})
	defer server.Close()
	err := api.SendDTMFToCall("123", &SendDTMFToCallData{DTMFOut: "1234"})
	if err != nil {
		t.Error("Failed call of SendDTMFToCall()")
		return
	}
}

func TestGetCallEvents(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/calls/123/events",
		Method:       http.MethodGet,
		ContentToSend: `[
		{
			"id": "{callEventId1}",
			"time": "2012-09-19T13:55:41.343Z",
			"name": "create"
		},
		{
			"id": "{callEventId2}",
			"time": "2012-09-19T13:55:45.583Z",
			"name": "answer"
		}]`}})
	defer server.Close()
	result, err := api.GetCallEvents("123")
	if err != nil {
		t.Error("Failed call of GetCallEvents()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetCallEventsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/events",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetCallEvents("123") })
}

func TestGetCallEvent(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/calls/123/events/456",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{callEventId1}",
			"time": "2012-09-19T13:55:41.343Z",
			"name": "create"
		}`}})
	defer server.Close()
	result, err := api.GetCallEvent("123", "456")
	if err != nil {
		t.Error("Failed call of GetCallEvent()")
		return
	}
	expect(t, result.ID, "{callEventId1}")
}

func TestGetCallEventFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/events/456",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetCallEvent("123", "456") })
}

func TestGetCallRecordings(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/calls/123/recordings",
		Method:       http.MethodGet,
		ContentToSend: `[
		{
			"id": "{callRecordingId1}"
		},
		{
			"id": "{callRecordingId2}"
		}]`}})
	defer server.Close()
	result, err := api.GetCallRecordings("123")
	if err != nil {
		t.Error("Failed call of GetCallRecordings()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetCallRecordingsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/recordings",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetCallRecordings("123") })
}

func TestGetCallTranscriptions(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/calls/123/transcriptions",
		Method:       http.MethodGet,
		ContentToSend: `[
		{
			"id": "{callTranscriptionId1}"
		},
		{
			"id": "{callTranscriptionId2}"
		}]`}})
	defer server.Close()
	result, err := api.GetCallTranscriptions("123")
	if err != nil {
		t.Error("Failed call of GetCallTranscriptions()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetCallTranscriptionsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/transcriptions",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetCallTranscriptions("123") })
}

func TestCreateGather(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/gather",
		Method:           http.MethodPost,
		EstimatedContent: `{"maxDigits":"5"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/calls/123/gather/456"}}})
	defer server.Close()
	id, err := api.CreateGather("123", &CreateGatherData{MaxDigits: 5})
	if err != nil {
		t.Error("Failed call of CreateGather()")
		return
	}
	expect(t, id, "456")
}

func TestCreateGatherFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/gather",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateGather("123", &CreateGatherData{MaxDigits: 5})
	})
}

func TestGetGather(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/calls/123/gather/456",
		Method:       http.MethodGet,
		ContentToSend: `{
		"id": "{gatherId}",
		"state": "completed",
		"reason": "max-digits",
		"createdTime": "2014-02-12T19:33:56Z",
		"completedTime": "2014-02-12T19:33:59Z",
		"call": "https://api.catapult.inetwork.com/v1/users/{userId}/calls/{callId}",
		"digits": "123"	}`}})
	defer server.Close()
	result, err := api.GetGather("123", "456")
	if err != nil {
		t.Error("Failed call of GetGather()")
		return
	}
	expect(t, result.ID, "{gatherId}")
}

func TestGetGatherFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/gather/456",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.GetGather("123", "456")
	})
}

func TestUpdateGather(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/gather/456",
		Method:           http.MethodPost,
		EstimatedContent: `{"state":"completed"}`}})
	defer server.Close()
	err := api.UpdateGather("123", "456", &UpdateGatherData{State: "completed"})
	if err != nil {
		t.Error("Failed call of UpdateGather()")
		return
	}
}

func TestUpdateGatherFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/calls/123/gather/456",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	err := api.UpdateGather("123", "456", &UpdateGatherData{State: "completed"})
	if err == nil {
		t.Error("Should fail here")
		return
	}
}
