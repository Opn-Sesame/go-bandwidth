package bandwidth


// ReservePhoneNumber reserves a phone number
// It returns ID of created phone number or error
// example: api.ReservePhoneNumber("+1-phone-number-to-buy")
func (api *Client) ReservePhoneNumber(phoneNumber string, options ...map[string]interface{}) (string, error) {
	data := map[string]interface{}{	"number": phoneNumber }
	if len(options) > 0 {
		mergeMaps(data, options[0])
	}
	return api.CreatePhoneNumber(data)
}
