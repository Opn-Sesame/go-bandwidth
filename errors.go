package bandwidth

import (
	"net/http"
	"fmt"
)

const errorsPath = "errors"

// GetErrors returns list of errors
func (api *Client) GetErrors(query ...map[string]string) ([]map[string] interface{}, error){
	var options map[string]string
	if len(query) > 0{
		options = query[0]
	}
	result, _, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(errorsPath), options, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}


// GetError returns  error by id
func (api *Client) GetError(id string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(errorsPath), id))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}
