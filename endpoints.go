package bandwidth

import (
	"fmt"
	"net/http"
)

const endpointsPath = "endpoints"

// DomainEndpoint struct
type DomainEndpoint struct {
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	DomainID      string            `json:"domainId"`
	ApplicationID string            `json:"applicationId"`
	Enabled       bool              `json:"enabled,string"`
	SipURI        string            `json:"sipUri"`
	Credentials   map[string]string `json:"credentials"`
}

// GetDomainEndpoints returns list of all endpoints for a domain
func (api *Client) GetDomainEndpoints(id string) ([]*DomainEndpoint, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath), &[]*DomainEndpoint{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*DomainEndpoint)), nil
}

// CreateDomainEndpoint creates a new endpoint for a domain
func (api *Client) CreateDomainEndpoint(id string, data map[string]interface{}) (string, error) {
	_, headers, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath), nil, data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// GetDomainEndpoint returns   single enpoint for a domain
func (api *Client) GetDomainEndpoint(id string, endpointID string) (*DomainEndpoint, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath, endpointID), &DomainEndpoint{})
	if err != nil {
		return nil, err
	}
	return result.(*DomainEndpoint), nil
}

// DeleteDomainEndpoint removes a endpoint from domain
func (api *Client) DeleteDomainEndpoint(id string, endpointID string) error {
	_, _, err := api.makeRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath, endpointID))
	return err
}

// UpdateDomainEndpoint removes a endpoint from domain
func (api *Client) UpdateDomainEndpoint(id string, endpointID string, data map[string]interface{}) error {
	_, _, err := api.makeRequest(http.MethodPost, fmt.Sprintf("%s/%s/%s/%s", api.concatUserPath(domainsPath), id, endpointsPath, endpointID), nil, data)
	return err
}
