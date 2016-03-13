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
		"balance": "100",
		"accountType": "pre-pay"
		}`}})
	defer server.Close()
	result, err := api.GetAccount()
	if err != nil {
		t.Error("Failed call of GetAccount()")
		return
	}
	expect(t, result["balance"], "100")
	expect(t, result["accountType"], "pre-pay")
}

func TestGetAccountFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/account",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetAccount() })
}
