package bandwidth

import (
	"net/http"
	"fmt"
)

const accountPath = "account"

// GetAccount returns account information (balance, etc)
func (api *Client) GetAccount() (map[string] interface{}, error){
	result, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(accountPath))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}


// GetAccountTransactions returns transactions from the user's account
func (api *Client) GetAccountTransactions() ([]map[string] interface{}, error){
	result, err :=  api.makeRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(accountPath), "transactions"), nil, []interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

