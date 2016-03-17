package bandwidth

import (
	"net/http"
	"fmt"
	"net/url"
)

const numberInfoPath = "phoneNumbers/numberInfo"

// GetNumberInfo returns information fo given number
func (api *Client) GetNumberInfo(number string) (map[string]interface{}, error) {
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", numberInfoPath, url.QueryEscape(number)))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}
