package bandwidth

import (
	"net/http"
	"fmt"
	"net/url"
)

// AvailableNumberType is allowed number types
type AvailableNumberType string

const (
	// AvailableNumberTypeLocal is local number
	AvailableNumberTypeLocal AvailableNumberType = "local"

	// AvailableNumberTypeTollFree is toll free number
	AvailableNumberTypeTollFree AvailableNumberType = "tollFree"
)

const availableNumbersPath = "availableNumbers"


// GetAvailableNumbers looks for available numbers
func (api *Client) GetAvailableNumbers(numberType AvailableNumberType, query map[string]string) ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", availableNumbersPath, numberType), query, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// GetAndOrderAvailableNumbers looks for available numbers and orders them
func (api *Client) GetAndOrderAvailableNumbers(numberType AvailableNumberType, query map[string]string) ([]map[string] interface{}, error){
	q := make(url.Values)
	for key, value := range query {
		q[key] = []string{value}
	}
	path := fmt.Sprintf("%s/%s?%s", availableNumbersPath, numberType, q.Encode())
	result, _, err :=  api.makeRequest(http.MethodPost, path, nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	list := result.([]map[string] interface{})
	for _, item := range list {
		location := item["location"]
		if location != nil {
			item["id"] = getIDFromLocation(location.(string))
		}
	}
	return list, nil
}

