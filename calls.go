package bandwidth

import (
	"net/http"
	"fmt"
)

const callsPath = "calls"

// GetCalls returns list of previous calls that were made or received
func (api *Client) GetCalls() ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(callsPath), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// CreateCall creates an outbound phone call
func (api *Client) CreateCall(data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, api.concatUserPath(callsPath), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetCall returns information about a call that was made or received
func (api *Client) GetCall(id string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(callsPath), id))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

// UpdateCall manage an active phone call. E.g. Answer an incoming call, reject an incoming call, turn on / off recording, transfer, hang up
func (api *Client) UpdateCall(id string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s", api.concatUserPath(callsPath), id), data)
	return err
}

// PlayAudioToCall plays an audio or speak a sentence in a call
func (api *Client) PlayAudioToCall(id string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "audio"), data)
	return err
}

// SendDTMFToCall plays an audio or speak a sentence in a call
func (api *Client) SendDTMFToCall(id string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "dtmf"), data)
	return err
}


// GetCallEvents returns  the list of call events for a call
func (api *Client) GetCallEvents(id string) ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "events"), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// GetCallEvent returns information about one call event
func (api *Client) GetCallEvent(id string, eventID string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(callsPath), id, "events", eventID))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

// GetCallRecordings returns  all recordings related to the call
func (api *Client) GetCallRecordings(id string) ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "recordings"), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// GetCallTranscriptions returns  all transcriptions  related to the call
func (api *Client) GetCallTranscriptions (id string) ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "transcriptions"), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// CreateGather gathers the DTMF digits pressed in a call
func (api *Client) CreateGather(id string, data map[string] interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(callsPath), id, "gather"), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetGather returns the gather DTMF parameters and results of the call
func (api *Client) GetGather(id string, gatherID string) (map[string]interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(callsPath), id, "gather", gatherID))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

// UpdateGather updates call's gather data
func (api *Client) UpdateGather(id string, gatherID string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(callsPath), id, "gather", gatherID), data)
	return err
}

