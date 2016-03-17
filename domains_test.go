package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetDomains(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/domains",
		Method:       http.MethodGet,
		ContentToSend: `[{
			"id": "{domainId1}",
			"name": "domain1"
		}, {
			"id": "{domainId2}",
			"name": "domain2"
		}]`}})
	defer server.Close()
	result, err := api.GetDomains()
	if err != nil {
		t.Error("Failed domain of GetDomains()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetDomainsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/domains",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetDomains() })
}

func TestCreateDomain(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/domains",
		Method:           http.MethodPost,
		EstimatedContent: `{"name":"domain"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/domains/123"}}})
	defer server.Close()
	id, err := api.CreateDomain(map[string]interface{}{"name": "domain"})
	if err != nil {
		t.Error("Failed domain of CreateDomain()")
		return
	}
	expect(t, id, "123")
}

func TestCreateDomainFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/domains",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateDomain(map[string]interface{}{"name": "domain"})
	})
}

func TestDeleteDomain(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/domains/123",
		Method:       http.MethodDelete}})
	defer server.Close()
	err := api.DeleteDomain("123")
	if err != nil {
		t.Error("Failed domain of DeleteDomain()")
		return
	}
}
