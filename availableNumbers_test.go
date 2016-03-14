package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetAvailableNumbers(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/availableNumbers/local?city=Cary&state=NC",
		Method:       http.MethodGet,
		ContentToSend: `[
		{
			"number": "{number1}",
			"nationalNumber": "{national_number1}",
			"patternMatch": "          2 9 ",
			"city": "CARY",
			"lata": "426",
			"rateCenter": "CARY",
			"state": "NC",
			"price": "0.60"
		},
		{
			"number": "{number2}",
			"nationalNumber": "{national_number2}",
			"patternMatch": "          2 9 ",
			"city": "CARY",
			"lata": "426",
			"rateCenter": "CARY",
			"state": "NC",
			"price": "0.60"
		}]`}})
	defer server.Close()
	result, err := api.GetAvailableNumbers(AvailableNumberTypeLocal, map[string]string{
		"state": "NC",
		"city": "Cary"})
	if err != nil {
		t.Errorf("Failed call of GetAvailableNumbers(): %s", err.Error())
		return
	}
	expect(t, len(result), 2)
}

func TestGetAvailableNumbersFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/availableNumbers/local?city=Cary&state=NC",
		Method:       http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func()(interface{}, error){ return api.GetAvailableNumbers(AvailableNumberTypeLocal, map[string]string{
		"state": "NC",
		"city": "Cary"})})
}

func TestGetAndOrderAvailableNumbers(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/availableNumbers/tollFree?city=Cary&state=NC",
		Method:       http.MethodPost,
		ContentToSend: `[
			{
				"number": "{number1}",
				"nationalNumber": "{national_number1}",
				"price": "0.60",
				"location": "https://.../v1/users/.../phoneNumbers/{numberId1}"
			},
			{
				"number": "{number2}",
				"nationalNumber": "{national_number2}",
				"price": "0.60",
				"location": "https://.../v1/users/.../phoneNumbers/{numberId2}"
			}
		]`}})
	defer server.Close()
	result, err := api.GetAndOrderAvailableNumbers(AvailableNumberTypeTollFree, map[string]string{
		"state": "NC",
		"city": "Cary"})
	if err != nil {
		t.Errorf("Failed call of GetAndOrderAvailableNumbers(): %s", err.Error())
		return
	}
	expect(t, len(result), 2)
	expect(t, result[0]["id"], "{numberId1}")
	expect(t, result[1]["id"], "{numberId2}")
}

func TestGetAndOrderAvailableNumbersFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/availableNumbers/tollFree?city=Cary&state=NC",
		Method:       http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func()(interface{}, error){ return api.GetAndOrderAvailableNumbers(AvailableNumberTypeTollFree, map[string]string{
		"state": "NC",
		"city": "Cary"})})
}

