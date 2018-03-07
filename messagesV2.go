package bandwidth

import (
	"net/http"
	"time"
)

// CreateMessageDataV2 struct
type CreateMessageDataV2 struct {
	From          string      `json:"from,omitempty"`
	To            interface{} `json:"to,omitempty"`
	Text          string      `json:"text,omitempty"`
	Media         []string    `json:"media,omitempty"`
	ApplicationID string      `json:"applicationId,omitempty"`
	Tag           string      `json:"tag,omitempty"`
}

// CreateMessageResultV2 stores status of sent message
type CreateMessageResultV2 struct {
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
func (api *Client) CreateMessageV2(data *CreateMessageDataV2) (*CreateMessageResultV2, error) {
	result, _, err := api.makeRequestV2(http.MethodPost, api.concatUserPath(messagesPath), &CreateMessageResultV2{}, data)
	if err != nil {
		return nil, err
	}
	return result.(*CreateMessageResultV2), nil
}
