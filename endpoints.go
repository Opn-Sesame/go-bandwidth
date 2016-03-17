package bandwidth

import (
	"net/http"
	"fmt"
)

const endpointsPath = "endpoints"

// GetDomainEndpoints returns list of all endpoints for a domain
func (api *Client) GetDomainEndpoints(id string) ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// CreateDomainEndpoint creates a new endpoint for a domain
func (api *Client) CreateDomainEndpoint(id string, data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetDomainEndpoint returns   single enpoint for a domain
func (api *Client) GetDomainEndpoint(id string, endpointID string) (map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath, endpointID))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

// DeleteDomainEndpoint removes a endpoint from domain
func (api *Client) DeleteDomainEndpoint(id string, endpointID string) error{
	_, _, err :=  api.makeRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath, endpointID))
	return err
}

// UpdateDomainEndpoint removes a endpoint from domain
func (api *Client) UpdateDomainEndpoint(id string, endpointID string, data map[string]interface{}) error{
	_, _, err :=  api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath, endpointID), data)
	return err
}
