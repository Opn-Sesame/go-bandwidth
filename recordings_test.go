package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetRecordings(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/recordings",
		Method:       http.MethodGet,
		ContentToSend: `[{
			"id": "{recordingId1}",
			"media": "recording1"
		}, {
			"id": "{recordingId2}",
			"media": "recording2"
		}]`}})
	defer server.Close()
	result, err := api.GetRecordings()
	if err != nil {
		t.Error("Failed call of GetRecordings()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetRecordingsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/recordings",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetRecordings() })
}

func TestGetRecording(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/recordings/123",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{recordingId1}",
			"media": "recording1"
		}`}})
	defer server.Close()
	result, err := api.GetRecording("123")
	if err != nil {
		t.Error("Failed call of GetRecording()")
		return
	}
	expect(t, result["media"], "recording1")
}

func TestGetRecordingFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/recordings/123",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetRecording("123") })
}
