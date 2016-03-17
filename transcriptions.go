package bandwidth

import (
	"net/http"
	"fmt"
)

const transcriptionsPath = "transcriptions"

// GetRecordingTranscriptions returns list of all transcriptions for a recording
func (api *Client) GetRecordingTranscriptions(id string) ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(recordingsPath), id, transcriptionsPath), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// CreateRecordingTranscription creates a new transcription for a recording
func (api *Client) CreateRecordingTranscription(id string, data ...map[string]interface{}) (string, error){
	item := map[string]interface{}{}
	if len(data) > 0 {
		item = data[0]
	}
	_, headers, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(recordingsPath), id, transcriptionsPath), item)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetRecordingTranscription returns   single enpoint for a recording
func (api *Client) GetRecordingTranscription(id string, transcriptionID string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(recordingsPath), id, transcriptionsPath, transcriptionID))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

