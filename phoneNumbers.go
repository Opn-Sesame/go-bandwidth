package bandwidth

import (
	"fmt"
	"net/http"
	"net/url"
)

const phoneNumbersPath = "phoneNumbers"

// PhoneNumber struct
type PhoneNumber struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Number         string  `json:"number"`
	NationalNumber string  `json:"nationalNumber"`
	City           string  `json:"city"`
	State          string  `json:"state"`
	ApplicationID  string  `json:"applicationId"`
	FallbackNumber string  `json:"fallbackNumber"`
	CreatedTime    string  `json:"createdTime"`
	NumberState    string  `json:"numberState"`
	Price          float64 `json:"price,string"`
}

// GetPhoneNumbers returns a list of your numbers
func (api *Client) GetPhoneNumbers() ([]*PhoneNumber, error) {
	result, _, err := api.makeRequest(http.MethodGet, api.concatUserPath(phoneNumbersPath), &[]*PhoneNumber{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*PhoneNumber)), nil
}

// CreatePhoneNumber creates a new phone number
func (api *Client) CreatePhoneNumber(data map[string]interface{}) (string, error) {
	_, headers, err := api.makeRequest(http.MethodPost, api.concatUserPath(phoneNumbersPath), nil, data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetPhoneNumber returns information for phone number by id or number
func (api *Client) GetPhoneNumber(idOrNumber string) (*PhoneNumber, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(phoneNumbersPath), url.QueryEscape(idOrNumber)), &PhoneNumber{})
	if err != nil {
		return nil, err
	}
	return result.(*PhoneNumber), nil
}

// UpdatePhoneNumber makes changes to your number
func (api *Client) UpdatePhoneNumber(idOrNumber string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s", api.concatUserPath(phoneNumbersPath), url.QueryEscape(idOrNumber)), nil, data)
	return err
}

// DeletePhoneNumber removes a phone number
func (api *Client) DeletePhoneNumber(id string) error {
	_, _, err := api.makeRequest(http.MethodDelete, fmt.Sprintf("%s/%s", api.concatUserPath(phoneNumbersPath), id))
	return err
}
