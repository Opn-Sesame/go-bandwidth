package bandwidth

import (
	"net/http"
	"fmt"
	"net/url"
)

const phoneNumbersPath = "phoneNumbers"

// GetPhoneNumbers returns a list of your numbers
func (api *Client) GetPhoneNumbers() ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(phoneNumbersPath), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// CreatePhoneNumber creates a new phone number
func (api *Client) CreatePhoneNumber(data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, api.concatUserPath(phoneNumbersPath), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetPhoneNumber returns information for phone number by id or number
func (api *Client) GetPhoneNumber(idOrNumber string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(phoneNumbersPath), url.QueryEscape(idOrNumber)))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

// UpdatePhoneNumber makes changes to your number
func (api *Client) UpdatePhoneNumber(idOrNumber string, data map[string]interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s", api.concatUserPath(phoneNumbersPath), url.QueryEscape(idOrNumber)), data)
	return err
}

// DeletePhoneNumber removes a phone number
func (api *Client) DeletePhoneNumber(id string) error{
	_, _, err :=  api.makeRequest(http.MethodDelete, fmt.Sprintf("%s/%s", api.concatUserPath(phoneNumbersPath), id))
	return err
}
