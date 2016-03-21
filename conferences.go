package bandwidth

import (
	"fmt"
	"net/http"
)

const conferencesPath = "conferences"

// Conference struct
type Conference struct {
	ID              string `json:"id"`
	State           string `json:"state"`
	From            string `json:"from"`
	CreatedTime     string `json:"createdTime"`
	ActiveMembers   int    `json:"activeMembers"`
	CallbackURL     string `json:"callbackUrl"`
	CallbackTimeout int    `json:"callbackTimeout,string"`
	FallbackURL     string `json:"fallbackUrl"`
	Hold            bool   `json:"hold,string"`
	Mute            bool   `json:"mute,string"`
}

// CreateConference creates a new conference
// It returns ID of creeated conference
func (api *Client) CreateConference(data map[string]interface{}) (string, error) {
	_, headers, err := api.makeRequest(http.MethodPost, api.concatUserPath(conferencesPath), nil, data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetConference returns information about a conference
//It return Conference instance for found conference or error
func (api *Client) GetConference(id string) (*Conference, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(conferencesPath), id), &Conference{})
	if err != nil {
		return nil, err
	}
	return result.(*Conference), nil
}

// UpdateConference manage an active phone conference. E.g. Answer an incoming conference, reject an incoming conference, turn on / off recording, transfer, hang up
// It returns error object
func (api *Client) UpdateConference(id string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s", api.concatUserPath(conferencesPath), id), nil, data)
	return err
}

// PlayAudioToConference plays an audio or speak a sentence in a conference
// It returns error object
func (api *Client) PlayAudioToConference(id string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(conferencesPath), id, "audio"), nil, data)
	return err
}

// ConferenceMember struct
type ConferenceMember struct {
	ID          string `json:"id"`
	State       string `json:"state"`
	AddedTime   string `json:"addedTime"`
	RemovedTime string `json:"removedTime"`
	Hold        bool   `json:"hold,string"`
	Mute        bool   `json:"mute,string"`
	JoinTone    bool   `json:"joinTone,string"`
	LeavingTone bool   `json:"leavingTone,string"`
}

// CreateConferenceMember creates a new conference member
// It returns ID of created member
func (api *Client) CreateConferenceMember(id string, data map[string]interface{}) (string, error) {
	_, headers, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(conferencesPath), id, "members"), nil, data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetConferenceMembers returns  the list of conference members
// It returns list of ConferenceMember or error
func (api *Client) GetConferenceMembers(id string) ([]*ConferenceMember, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(conferencesPath), id, "members"), &[]*ConferenceMember{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*ConferenceMember)), nil
}

// GetConferenceMember returns information about one conference member
// It returns ConferenceMember instance for found instance or error
func (api *Client) GetConferenceMember(id string, memberID string) (*ConferenceMember, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(conferencesPath), id, "members", memberID), &ConferenceMember{})
	if err != nil {
		return nil, err
	}
	return result.(*ConferenceMember), nil
}

// UpdateConferenceMember updates a conference member
// It returns error object
func (api *Client) UpdateConferenceMember(id string, memberID string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(conferencesPath), id, "members", memberID), nil, data)
	return err
}

// PlayAudioToConferenceMember plays an audio or speak a sentence to a conference member
// It returns error object
func (api *Client) PlayAudioToConferenceMember(id string, memberID string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s/%s", api.concatUserPath(conferencesPath), id, "members", memberID, "audio"), nil, data)
	return err
}
