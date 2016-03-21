package bandwidth

import (
	"fmt"
	"net/http"
)

const transcriptionsPath = "transcriptions"

// Transcription struct
type Transcription struct {
	ID                 string `json:"id"`
	ChargeableDuration int    `json:"chargeableDuration"`
	Text               string `json:"text"`
	TextSize           int    `json:"textSize"`
	TextURL            string `json:"textUrl"`
	Time               string `json:"time"`
}

// GetRecordingTranscriptions returns list of all transcriptions for a recording
func (api *Client) GetRecordingTranscriptions(id string) ([]*Transcription, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(recordingsPath), id, transcriptionsPath), &[]*Transcription{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Transcription)), nil
}

// CreateRecordingTranscription creates a new transcription for a recording
func (api *Client) CreateRecordingTranscription(id string, data ...map[string]interface{}) (string, error) {
	item := map[string]interface{}{}
	if len(data) > 0 {
		item = data[0]
	}
	_, headers, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(recordingsPath), id, transcriptionsPath), nil, item)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetRecordingTranscription returns   single enpoint for a recording
func (api *Client) GetRecordingTranscription(id string, transcriptionID string) (*Transcription, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(recordingsPath), id, transcriptionsPath, transcriptionID), &Transcription{})
	if err != nil {
		return nil, err
	}
	return result.(*Transcription), nil
}
