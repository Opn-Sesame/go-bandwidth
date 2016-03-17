package bandwidth

import (
	"net/http"
	"fmt"
)

const conferencesPath = "conferences"

// CreateConference creates a new conference
func (api *Client) CreateConference(data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, api.concatUserPath(conferencesPath), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetConference returns information about a conference
func (api *Client) GetConference(id string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(conferencesPath), id))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

// UpdateConference manage an active phone conference. E.g. Answer an incoming conference, reject an incoming conference, turn on / off recording, transfer, hang up
func (api *Client) UpdateConference(id string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s", api.concatUserPath(conferencesPath), id), data)
	return err
}

// PlayAudioToConference plays an audio or speak a sentence in a conference
func (api *Client) PlayAudioToConference(id string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(conferencesPath), id, "audio"), data)
	return err
}


// CreateConferenceMember creates a new conference member
func (api *Client) CreateConferenceMember(id string, data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(conferencesPath), id, "members"), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetConferenceMembers returns  the list of conference members
func (api *Client) GetConferenceMembers(id string) ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(conferencesPath), id, "members"), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// GetConferenceMember returns information about one conference member
func (api *Client) GetConferenceMember(id string, memberID string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(conferencesPath), id, "members", memberID))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

// UpdateConferenceMember updates a conference member
func (api *Client) UpdateConferenceMember(id string, memberID string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(conferencesPath), id, "members", memberID), data)
	return err
}


// PlayAudioToConferenceMember plays an audio or speak a sentence to a conference member
func (api *Client) PlayAudioToConferenceMember(id string, memberID string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s/%s", api.concatUserPath(conferencesPath), id, "members", memberID, "audio"), data)
	return err
}
