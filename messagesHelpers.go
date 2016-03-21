package bandwidth

// SendMessageTo sends a SMS/MMS to given number
// It returns ID of created message or error
func (api *Client) SendMessageTo(fromNumber string, toNumber string, text string, options ...map[string]interface{}) (string, error){
	data := map[string]interface{}{
		"from": fromNumber,
		"to": toNumber,
		"text": text }
	if len(options) > 0 {
		mergeMaps(data, options[0])
	}
	return api.CreateMessage(data)
}
