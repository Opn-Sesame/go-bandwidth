package bandwidth

import (
	"fmt"
	"net/http"
)

const bridgesPath = "bridges"

// Bridge struct
type Bridge struct {
	ID            string `json:"id"`
	State         string `json:"state"`
	BridgeAudio   bool   `json:"bridgeAudio,string"`
	Calls         string `json:"calls"`
	CreatedTime   string `json:"createdTime"`
	ActivatedTime string `json:"activatedTime"`
	CompletedTime string `json:"completedTime"`
}

// GetBridges returns list of previous bridges
func (api *Client) GetBridges() ([]*Bridge, error) {
	result, _, err := api.makeRequest(http.MethodGet, api.concatUserPath(bridgesPath), &[]*Bridge{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Bridge)), nil
}

// CreateBridge creates a bridge
func (api *Client) CreateBridge(data map[string]interface{}) (string, error) {
	_, headers, err := api.makeRequest(http.MethodPost, api.concatUserPath(bridgesPath), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetBridge returns a bridge
func (api *Client) GetBridge(id string) (*Bridge, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(bridgesPath), id), &Bridge{})
	if err != nil {
		return nil, err
	}
	return result.(*Bridge), nil
}

// UpdateBridge adds one or two calls in a bridge and also puts the bridge on hold/unhold
func (api *Client) UpdateBridge(id string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s", api.concatUserPath(bridgesPath), id), data)
	return err
}

// PlayAudioToBridge plays an audio or speak a sentence in a bridge
func (api *Client) PlayAudioToBridge(id string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(bridgesPath), id, "audio"), data)
	return err
}

// GetBridgeCalls returns bridge's calls
func (api *Client) GetBridgeCalls(id string) ([]*Call, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(bridgesPath), id, "calls"), &[]*Call{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Call)), nil
}
