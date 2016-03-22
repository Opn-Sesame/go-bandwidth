package bandwidth

import (
	"fmt"
	"net/http"
)

const messagesPath = "messages"

// Message struct
type Message struct {
	ID                  string   `json:"id"`
	From                string   `json:"from"`
	To                  string   `json:"to"`
	Direction           string   `json:"direction"`
	Text                string   `json:"text"`
	Media               []string `json:"media"`
	State               string   `json:"state"`
	Time                string   `json:"time"`
	CallbackURL         string   `json:"callbackUrl"`
	ReceiptRequested    string   `json:"receiptRequested"`
	DeliveryState       string   `json:"deliveryState"`
	DeliveryCode        string   `json:"deliveryCode"`
	DeliveryDescription string   `json:"deliveryDescription"`
}

// GetMessages returns list of all messages
// It returns list of Message instances or error
func (api *Client) GetMessages(query ...map[string]string) ([]*Message, error) {
	var options map[string]string
	if len(query) > 0 {
		options = query[0]
	}
	result, _, err := api.makeRequest(http.MethodGet, api.concatUserPath(messagesPath), &[]*Message{}, options)
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Message)), nil
}

// CreateMessage sends a message (SMS/MMS)
// It returns ID of created message or error
func (api *Client) CreateMessage(data interface{}) (string, error) {
	_, headers, err := api.makeRequest(http.MethodPost, api.concatUserPath(messagesPath), nil, data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetMessage returns a single message
// It returns Message instance or error
func (api *Client) GetMessage(id string) (*Message, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(messagesPath), id), &Message{})
	if err != nil {
		return nil, err
	}
	return result.(*Message), nil
}
