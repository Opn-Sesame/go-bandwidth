package bandwidth

import (
	"fmt"
	"net/http"
)

const callsPath = "calls"

// Call struct
type Call struct {
	ID                   string `json:"id"`
	ActiveTime           string `json:"activeTime"`
	ChargeableDuration   int    `json:"chargeableDuration"`
	Direction            string `json:"direction"`
	Events               string `json:"events"`
	EndTime              string `json:"endTime"`
	From                 string `json:"from"`
	RecordingFileFormat  string `json:"recordingFileFormat"`
	RecordingEnabled     bool   `json:"recordingEnabled"`
	StartTime            string `json:"startTime"`
	State                string `json:"state"`
	To                   string `json:"to"`
	TranscriptionEnabled bool   `json:"transcriptionEnabled"`
}

// GetCalls returns list of previous calls that were made or received
func (api *Client) GetCalls() ([]*Call, error) {
	result, _, err := api.makeRequest(http.MethodGet, api.concatUserPath(callsPath), &[]*Call{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Call)), nil
}

// CreateCall creates an outbound phone call
func (api *Client) CreateCall(data map[string]interface{}) (string, error) {
	_, headers, err := api.makeRequest(http.MethodPost, api.concatUserPath(callsPath), nil, data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetCall returns information about a call that was made or received
func (api *Client) GetCall(id string) (*Call, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(callsPath), id), &Call{})
	if err != nil {
		return nil, err
	}
	return result.(*Call), nil
}

// UpdateCall manage an active phone call. E.g. Answer an incoming call, reject an incoming call, turn on / off recording, transfer, hang up
func (api *Client) UpdateCall(id string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s", api.concatUserPath(callsPath), id), nil, data)
	return err
}

// PlayAudioToCall plays an audio or speak a sentence in a call
func (api *Client) PlayAudioToCall(id string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "audio"), nil, data)
	return err
}

// SendDTMFToCall plays an audio or speak a sentence in a call
func (api *Client) SendDTMFToCall(id string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "dtmf"), nil, data)
	return err
}

// CallEvent struct
type CallEvent struct {
	ID   string `json:"id"`
	Time string `json:"time"`
	Name string `json:"name"`
}

// GetCallEvents returns  the list of call events for a call
func (api *Client) GetCallEvents(id string) ([]*CallEvent, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "events"), nil, &[]*CallEvent{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*CallEvent)), nil
}

// GetCallEvent returns information about one call event
func (api *Client) GetCallEvent(id string, eventID string) (*CallEvent, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(callsPath), id, "events", eventID), &CallEvent{})
	if err != nil {
		return nil, err
	}
	return result.(*CallEvent), nil
}

// GetCallRecordings returns  all recordings related to the call
func (api *Client) GetCallRecordings(id string) ([]*Recording, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "recordings"), nil, &[]*Recording{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Recording)), nil
}

// GetCallTranscriptions returns  all transcriptions  related to the call
func (api *Client) GetCallTranscriptions(id string) ([]*Transcription, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "transcriptions"), nil, &[]*Transcription{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Transcription)), nil
}

// CreateGather gathers the DTMF digits pressed in a call
func (api *Client) CreateGather(id string, data map[string]interface{}) (string, error) {
	_, headers, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "gather"), nil, data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// Gather struct
type Gather struct {
	ID            string `json:"id"`
	State         string `json:"state"`
	Reason        string `json:"reason"`
	CreatedTime   string `json:"createdTime"`
	CompletedTime string `json:"completedTime"`
	Digits        string `json:"digits"`
}

// GetGather returns the gather DTMF parameters and results of the call
func (api *Client) GetGather(id string, gatherID string) (*Gather, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(callsPath), id, "gather", gatherID), &Gather{})
	if err != nil {
		return nil, err
	}
	return result.(*Gather), nil
}

// UpdateGather updates call's gather data
func (api *Client) UpdateGather(id string, gatherID string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(callsPath), id, "gather", gatherID), nil, data)
	return err
}
