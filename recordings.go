package bandwidth

import (
	"fmt"
	"net/http"
)

const recordingsPath = "recordings"

// Recording struct
type Recording struct {
	ID        string `json:"id"`
	EndTime   string `json:"endTime"`
	Media     string `json:"media"`
	Call      string `json:"call"`
	StartTime string `json:"startTime"`
	State     string `json:"state"`
}

// GetRecordings returns  a list of the calls recordings
func (api *Client) GetRecordings() ([]*Recording, error) {
	result, _, err := api.makeRequest(http.MethodGet, api.concatUserPath(recordingsPath), &[]*Recording{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Recording)), nil
}

// GetRecording returns  a single call recording
func (api *Client) GetRecording(id string) (*Recording, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(recordingsPath), id), &Recording{})
	if err != nil {
		return nil, err
	}
	return result.(*Recording), nil
}
