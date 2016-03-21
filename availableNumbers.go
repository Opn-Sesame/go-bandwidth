package bandwidth

import (
	"fmt"
	"net/http"
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

// AvailableNumber struct
type AvailableNumber struct {
	Number         string  `json:"number"`
	NationalNumber string  `json:"nationalNumber"`
	City           string  `json:"city"`
	RateCenter     string  `json:"rateCenter"`
	State          string  `json:"state"`
	Price          float64 `json:"price,string"`
}

// GetAvailableNumbers looks for available numbers
func (api *Client) GetAvailableNumbers(numberType AvailableNumberType, query map[string]string) ([]*AvailableNumber, error) {
	result, _, err := api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", availableNumbersPath, numberType), &[]*AvailableNumber{}, query)
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*AvailableNumber)), nil
}

// OrderedNumber struct
type OrderedNumber struct {
	Number         string  `json:"number"`
	NationalNumber string  `json:"nationalNumber"`
	Price          float64 `json:"price,string"`
	Location       string  `json:"Location"`
	ID             string  `json:"-"`
}

// GetAndOrderAvailableNumbers looks for available numbers and orders them
func (api *Client) GetAndOrderAvailableNumbers(numberType AvailableNumberType, query map[string]string) ([]*OrderedNumber, error) {
	q := make(url.Values)
	for key, value := range query {
		q[key] = []string{value}
	}
	path := fmt.Sprintf("%s/%s?%s", availableNumbersPath, numberType, q.Encode())
	result, _, err := api.makeRequest(http.MethodPost, path, &[]*OrderedNumber{})
	if err != nil {
		return nil, err
	}
	list := *(result.(*[]*OrderedNumber))
	for _, item := range list {
		if item.Location != "" {
			item.ID = getIDFromLocation(item.Location)
		}
	}
	return list, nil
}
