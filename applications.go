package bandwidth

import (
	"net/http"
	"fmt"
)

const applicationsPath = "applications"

// GetApplications returns list of user's applications
func (api *Client) GetApplications(query ...map[string]string) ([]map[string] interface{}, error){
	var options map[string]string
	if len(query) > 0{
		options = query[0]
	}
	result, _, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(applicationsPath), options, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// CreateApplication creates an application that can handle calls and messages for one of your phone number. Many phone numbers can share an application.
func (api *Client) CreateApplication(data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, api.concatUserPath(applicationsPath), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetApplication returns an user's application
func (api *Client) GetApplication(id string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(applicationsPath), id))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

// UpdateApplication makes changes to an application
func (api *Client) UpdateApplication(id string, data map[string] interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s", api.concatUserPath(applicationsPath), id), data)
	return err
}

// DeleteApplication permanently deletes an application
func (api *Client) DeleteApplication(id string) error{
	_, _, err :=  api.makeRequest(http.MethodDelete, fmt.Sprintf("%s/%s", api.concatUserPath(applicationsPath), id))
	return err
}
