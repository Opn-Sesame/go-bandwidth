package bandwidth

import (
	"net/http"
	"time"
)

// CreateMessage struct
type CreateMessage struct {
	From          string      `json:"from,omitempty"`
	To            interface{} `json:"to,omitempty"`
	Text          string      `json:"text,omitempty"`
	Media         []string    `json:"media,omitempty"`
	ApplicationID string      `json:"applicationId,omitempty"`
	Tag           string      `json:"tag,omitempty"`
}

// CreateMessageResponse stores status of sent message
type CreateMessageResponse struct {
	ID            string      `json:"id"`
	Time          *time.Time  `json:"time,string"`
	From          string      `json:"from"`
	To            interface{} `json:"to"`
	Text          string      `json:"text"`
	Media         []string    `json:"media"`
	ApplicationID string      `json:"applicationId"`
	Tag           string      `json:"tag"`
	Direction     string      `json:"direction"`
	SegmentCount  int32       `json:"segmentCount"`
}

// CreateMessageV2 sends a message (SMS/MMS)
func (c *Client) CreateMessage(data *CreateMessage) (*CreateMessageResponse, error) {
	result, _, err := c.makeMessagingRequest(http.MethodPost, c.MessagingEndpoint, &CreateMessageResponse{}, data)
	if err != nil {
		return nil, err
	}
	return result.(*CreateMessageResponse), nil
}
