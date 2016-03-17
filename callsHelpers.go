package bandwidth

func mergeMaps(src, dst map[string]interface{}){
	if dst == nil {
		dst = map[string]interface{}{}
	}
	for k, v := range dst {
		src[k] = v
	}
}

// CallTo creates call to given phone number
func (api *Client) CallTo(fromNumber string, toNumber string, options ...map[string]interface{}) (string, error){
	data := map[string]interface{}{
		"from": fromNumber,
		"to": toNumber }
	if len(options) > 0 {
		mergeMaps(data, options[0])
	}
	return api.CreateCall(data)
}

// AnswerIncomingCall  answers an incoming call
func (api *Client) AnswerIncomingCall(id string) error{
	return api.UpdateCall(id,  map[string]interface{}{"state": "active"})
}

// RejectIncomingCall  answers an incoming call
func (api *Client) RejectIncomingCall(id string) error{
	return api.UpdateCall(id,  map[string]interface{}{"state": "rejected"})
}

// HangUpCall  hangs up the call
func (api *Client) HangUpCall(id string) error{
	return api.UpdateCall(id,  map[string]interface{}{"state": "completed"})
}

// SetCallRecodingEnabled  hangs up the call
func (api *Client) SetCallRecodingEnabled(id string, enabled bool) error{
	return api.UpdateCall(id,  map[string]interface{}{"recordingEnabled": enabled})
}

// TransferCallTo  transfers call to another number
func (api *Client) TransferCallTo(id string, transferToNumber string, options ...map[string]interface{}) error{
	data := map[string]interface{}{
		"state": "transferring",
		"transferTo": transferToNumber }
	if len(options) > 0 {
		mergeMaps(data, options[0])
	}
	return api.UpdateCall(id,  data)
}


// StopGather stops call's gather
func (api *Client) StopGather(id string, gatherID string) error{
	return api.UpdateGather(id, gatherID, map[string]interface{}{"state": "completed"})
}

// SendDTMFCharactersToCall sends some dtmf chanracters to call
func (api *Client) SendDTMFCharactersToCall(id string, dtmfOut string) error{
	return api.SendDTMFToCall(id,  map[string]interface{}{"dtmfOut": dtmfOut})
}
