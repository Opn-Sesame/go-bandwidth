package bandwidth

import (
	"net/http"
	"fmt"
)

const bridgesPath = "bridges"

// GetBridges returns list of previous bridges
func (api *Client) GetBridges() ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(bridgesPath), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// CreateBridge creates a bridge
func (api *Client) CreateBridge(data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, api.concatUserPath(bridgesPath), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetBridge returns a bridge
func (api *Client) GetBridge(id string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(bridgesPath), id))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

// UpdateBridge adds one or two calls in a bridge and also puts the bridge on hold/unhold
func (api *Client) UpdateBridge(id string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s", api.concatUserPath(bridgesPath), id), data)
	return err
}

// PlayAudioToBridge plays an audio or speak a sentence in a bridge
func (api *Client) PlayAudioToBridge(id string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(bridgesPath), id, "audio"), data)
	return err
}

// GetBridgeCalls returns bridge's calls
func (api *Client) GetBridgeCalls(id string) ([]map[string]interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(bridgesPath), id, "calls"), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string]interface{}), nil
}
