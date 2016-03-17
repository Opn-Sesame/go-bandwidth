package bandwidth

import (
	"net/http"
	"fmt"
)

const messagesPath = "messages"

// GetMessages returns list of all messages
func (api *Client) GetMessages() ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(messagesPath), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// CreateMessage sends a message (SMS/MMS)
func (api *Client) CreateMessage(data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, api.concatUserPath(messagesPath), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetMessage returns a single message
func (api *Client) GetMessage(id string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(messagesPath), id))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}
