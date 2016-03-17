package bandwidth

import (
	"net/http"
	"fmt"
)

const recordingsPath = "recordings"

// GetRecordings returns  a list of the calls recordings
func (api *Client) GetRecordings() ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(recordingsPath), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// GetRecording returns  a single call recording
func (api *Client) GetRecording(id string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(recordingsPath), id))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}
