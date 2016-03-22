package bandwidth

import (
	"net/http"
	"testing"
)


func TestReservePhoneNumber(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/phoneNumbers",
		Method:           http.MethodPost,
		EstimatedContent: `{"number":"phoneNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/phoneNumbers/123"}}})
	defer server.Close()
	id, err := api.ReservePhoneNumber("phoneNumber")
	if err != nil {
		t.Error("Failed call of ReservePhoneNumber()")
		return
	}
	expect(t, id, "123")
}

func TestReservePhoneNumberWithOptions(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/phoneNumbers",
		Method:           http.MethodPost,
		EstimatedContent: `{"applicationId":"id","number":"phoneNumber"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/phoneNumbers/123"}}})
	defer server.Close()
	id, err := api.ReservePhoneNumber("phoneNumber", map[string]interface{}{"applicationId": "id"})
	if err != nil {
		t.Error("Failed call of ReservePhoneNumber()")
		return
	}
	expect(t, id, "123")
}
