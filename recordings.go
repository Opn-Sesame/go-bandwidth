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
// It returns list of Recording instances or error
func (api *Client) GetRecordings(query ...map[string]string) ([]*Recording, error) {
	var options map[string]string
	if len(query) > 0 {
		options = query[0]
	}
	result, _, err := api.makeRequest(http.MethodGet, api.concatUserPath(recordingsPath), &[]*Recording{}, options)
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Recording)), nil
}

// GetRecording returns  a single call recording
// It a Recording instance or error
func (api *Client) GetRecording(id string) (*Recording, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(recordingsPath), id), &Recording{})
	if err != nil {
		return nil, err
	}
	return result.(*Recording), nil
}
