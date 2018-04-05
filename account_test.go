package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetAccount(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/account",
		Method:       http.MethodGet,
		ContentToSend: `{
		"balance": 100,
		"accountType": "pre-pay"
		}`}})
	defer server.Close()
	result, err := api.GetAccount()
	if err != nil {
		t.Error("Failed call of GetAccount()")
		return
	}
	expect(t, result.Balance, 100.0)
	expect(t, result.AccountType, "pre-pay")
}

func TestGetAccountFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/account",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetAccount() })
}

func TestGetAccountTransactions(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/account/transactions",
		Method:       http.MethodGet,
		ContentToSend: `[
		{
			"id": "{transactionId1}",
			"time": "2013-02-21T13:39:09.122Z",
			"amount": "0.00750",
			"type": "charge",
			"units": 1,
			"productType": "sms-out",
			"number": "{number}"
		},
		{
			"id": "{transactionId2}",
			"time": "2013-02-21T13:37:42.079Z",
			"amount": "0.00750",
			"type": "charge",
			"units": 1,
			"productType": "sms-out",
			"number": "{number}"
		}
		]`}})
	defer server.Close()
	result, err := api.GetAccountTransactions()
	if err != nil {
		t.Error("Failed call of GetAccountTransactions()")
		return
	}
	expect(t, len(result), 2)
	expect(t, result[0].ID, "{transactionId1}")
	expect(t, result[1].ID, "{transactionId2}")
}

func TestGetAccountTransactionsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/account/transactions",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetAccountTransactions() })
}
