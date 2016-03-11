package bandwidth

const accountPath = "account"

// GetAccount returns account data (balance, etc)
func (api *Client) GetAccount() (map[string] interface{}, error){
	result, err :=  api.makeRequest("GET", api.concatUserPath(accountPath))
	if err != nil {
		return nil, err
	}
	return result.(map[string] interface{}), nil
}

