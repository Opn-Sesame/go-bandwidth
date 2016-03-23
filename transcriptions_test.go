package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetRecordingTranscriptions(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/recordings/123/transcriptions",
		Method:       http.MethodGet,
		ContentToSend: `[{
			"id": "{transcriptionId1}",
			"text": "transcription1"
		}, {
			"id": "{transcriptionId2}",
			"text": "transcription2"
		}]`}})
	defer server.Close()
	result, err := api.GetRecordingTranscriptions("123")
	if err != nil {
		t.Error("Failed call of GetRecordingTranscriptions()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetRecordingTranscriptionsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/recordings/123/transcriptions",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetRecordingTranscriptions("123") })
}

func TestCreateRecordingTranscription(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/recordings/123/transcriptions",
		Method:           http.MethodPost,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/recording/123/transcriptions/456"}}})
	defer server.Close()
	id, err := api.CreateRecordingTranscription("123")
	if err != nil {
		t.Error("Failed call of CreateRecordingTranscription()")
		return
	}
	expect(t, id, "456")
}

func TestCreateRecordingTranscriptionFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:    "/v1/users/userId/recordings/123/transcriptions",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateRecordingTranscription("123")
	})
}


func TestGetRecordingTranscription(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/recordings/123/transcriptions/456",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{transcriptionId2}",
			"text": "transcription2"
		}`}})
	defer server.Close()
	result, err := api.GetRecordingTranscription("123", "456")
	if err != nil {
		t.Error("Failed call of GetRecordingTranscription()")
		return
	}
	expect(t, result.Text, "transcription2")
}

func TestGetRecordingTranscriptionFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/recordings/123/transcriptions/456",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetRecordingTranscription("123", "456") })
}
