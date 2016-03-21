package bandwidth

import (
	"net/http"
	"fmt"
)

const messagesPath = "messages"

// Message struct
type Message struct {
	ID       string         `json:"id"`
	Category string         `json:"category"`
	Time     string         `json:"time"`
	Code     string         `json:"code"`
	Details  []*ErrorDetail `json:"details"`
}

// GetMessages returns list of all messages
func (api *Client) GetMessages() ([]*Message, error){
	result, _, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(messagesPath), &[]*Message{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Message)), nil
}

// CreateMessage sends a message (SMS/MMS)
func (api *Client) CreateMessage(data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, api.concatUserPath(messagesPath), nil, data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetMessage returns a single message
func (api *Client) GetMessage(id string) (*Message, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(messagesPath), id), &Message{})
	if err != nil {
		return nil, err
	}
	return result.(*Message), nil
}
