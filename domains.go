package bandwidth

import (
	"fmt"
	"net/http"
)

const domainsPath = "domains"

// Domain struct
type Domain struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetDomains returns  a list of the domains that have been created
func (api *Client) GetDomains() ([]*Domain, error) {
	result, _, err := api.makeRequest(http.MethodGet, api.concatUserPath(domainsPath), nil, &[]*Domain{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*Domain)), nil
}

// CreateDomain creates a new domain
func (api *Client) CreateDomain(data map[string]interface{}) (string, error) {
	_, headers, err := api.makeRequest(http.MethodPost, api.concatUserPath(domainsPath), nil, data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// DeleteDomain removes a domain
func (api *Client) DeleteDomain(id string) error {
	_, _, err := api.makeRequest(http.MethodDelete, fmt.Sprintf("%s/%s", api.concatUserPath(domainsPath), id))
	return err
}
