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
func (api *Client) CreateMessageV2(data *CreateMessageDataV2, other ...string) (*CreateMessageResultV2, error) {
    var v2_endpoint = "https://messaging.bandwidth.com"
    if len(other) > 0 {
        v2_endpoint = other[0]
    }

    var currentAPIEndPoint = api.APIEndPoint
    api.APIEndPoint = v2_endpoint
	result, _, err := api.makeRequestV2(http.MethodPost, api.concatUserPath(messagesPath), &CreateMessageResultV2{}, data)
    api.APIEndPoint = currentAPIEndPoint
	if err != nil {
		return nil, err
	}
	return result.(*CreateMessageResultV2), nil
}
