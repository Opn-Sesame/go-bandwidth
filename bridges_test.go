package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetBridges(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/bridges",
		Method:       http.MethodGet,
		ContentToSend: `[
		{
			"id": "{bridgeId}",
			"state": "completed",
			"bridgeAudio": true,
			"calls":"https://.../v1/users/{userId}/bridges/{bridgeId}/calls",
			"createdTime": "2013-04-22T13:55:30.279Z",
			"activatedTime": "2013-04-22T13:55:30.280Z",
			"completedTime": "2013-04-22T13:56:30.122Z"
		},
		{
			"id": "{bridgeId}",
			"state": "completed",
			"bridgeAudio": true,
			"calls":"https://.../v1/users/{userId}/bridges/{bridgeId}/calls",
			"createdTime": "2013-04-22T13:58:30.121Z",
			"activatedTime": "2013-04-22T13:58:30.122Z",
			"completedTime": "2013-04-22T13:59:30.122Z"
		}
		]`}})
	defer server.Close()
	result, err := api.GetBridges()
	if err != nil {
		t.Error("Failed call of GetBridges()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetBridgesFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/bridges",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetBridges() })
}

func TestCreateBridge(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/bridges",
		Method:           http.MethodPost,
		EstimatedContent: `{"bridgeAudio":true,"callIds":["{callId1}","{callId2}"]}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/bridges/123"}}})
	defer server.Close()
	id, err := api.CreateBridge(&BridgeData{
		BridgeAudio: true,
		CallIDs:     []string{"{callId1}", "{callId2}"}})
	if err != nil {
		t.Error("Failed call of CreateBridge()")
		return
	}
	expect(t, id, "123")
}

func TestCreateBridgeFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/bridges",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateBridge(&BridgeData{
			BridgeAudio: true,
			CallIDs:     []string{"{callId1}", "{callId2}"}})
	})
}

func TestGetBridge(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/bridges/123",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{bridgeId}",
			"state": "completed",
			"bridgeAudio": true,
			"calls":"https://.../v1/users/{userId}/bridges/{bridgeId}/calls",
			"createdTime": "2013-04-22T13:58:30.121Z",
			"activatedTime": "2013-04-22T13:58:30.122Z",
			"completedTime": "2013-04-22T13:59:30.122Z"
		}`}})
	defer server.Close()
	result, err := api.GetBridge("123")
	if err != nil {
		t.Error("Failed call of GetBridge()")
		return
	}
	expect(t, result.ID, "{bridgeId}")
}

func TestGetBridgeFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/bridges/123",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetBridge("123") })
}

func TestUpdateBridge(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/bridges/123",
		Method:           http.MethodPost,
		EstimatedContent: `{"callIds":["{callId1}","{callId2}"]}`}})
	defer server.Close()
	err := api.UpdateBridge("123", &BridgeData{
		CallIDs: []string{"{callId1}", "{callId2}"}})
	if err != nil {
		t.Error("Failed call of UpdateBridge()")
		return
	}
}

func TestPlayAudioToBridge(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/bridges/123/audio",
		Method:           http.MethodPost,
		EstimatedContent: `{"fileUrl":"file.mp3"}`}})
	defer server.Close()
	err := api.PlayAudioToBridge("123", &PlayAudioData{FileURL: "file.mp3"})
	if err != nil {
		t.Error("Failed call of PlayAudioToBridge()")
		return
	}
}

func TestGetBridgeCalls(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/bridges/123/calls",
		Method:       http.MethodGet,
		ContentToSend: `[
			{
				"activeTime": "2013-05-22T19:49:39Z",
				"direction": "out",
				"from": "{fromNumber}",
				"id": "{callId1}",
				"bridgeId": "{bridgeId}",
				"startTime": "2013-05-22T19:49:35Z",
				"state": "active",
				"to": "{toNumber1}",
				"recordingEnabled": false,
				"events": "https://api.catapult.inetwork.com/v1/users/{userId}/calls/{callId1}/events",
				"bridge": "https://api.catapult.inetwork.com/v1/users/{userId}/bridges/{bridgeId}"
			},
			{
				"activeTime": "2013-05-22T19:50:16Z",
				"direction": "out",
				"from": "{fromNumber}",
				"id": "{callId2}",
				"bridgeId": "{bridgeId}",
				"startTime": "2013-05-22T19:50:16Z",
				"state": "active",
				"to": "{toNumber2}",
				"recordingEnabled": false,
				"events": "https://api.catapult.inetwork.com/v1/users/{userId}/calls/{callId2}/events",
				"bridge": "https://api.catapult.inetwork.com/v1/users/{userId}/bridges/{bridgeId}"
			}
		]`}})
	defer server.Close()
	result, err := api.GetBridgeCalls("123")
	if err != nil {
		t.Error("Failed call of GetBridgeCalls()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetBridgeCallsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/bridges/123/calls",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetBridgeCalls("123") })
}
