package bandwidth

import (
	"fmt"
	"net/http"
)

const errorsPath = "errors"

// Error struct
type Error struct {
	ID       string         `json:"id"`
	Category string         `json:"category"`
	Time     string         `json:"time"`
	Code     string         `json:"code"`
	Details  []*ErrorDetail `json:"details"`
}

// ErrorDetail struct
type ErrorDetail struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GetErrors returns list of errors
func (api *Client) GetErrors(query ...map[string]string) ([]*Error, error) {
	var options map[string]string
	if len(query) > 0 {
		options = query[0]
	}
	result, _, err := api.makeRequest(http.MethodGet, api.concatUserPath(errorsPath), &[]*Error{}, options)
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Error)), nil
}

// GetError returns  error by id
func (api *Client) GetError(id string) (*Error, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(errorsPath), id), &Error{})
	if err != nil {
		return nil, err
	}
	return result.(*Error), nil
}
