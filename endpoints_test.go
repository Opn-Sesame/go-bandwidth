package bandwidth

import (
	"net/http"
	"testing"
)

func TestGetDomainEndpoints(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/domains/123/endpoints",
		Method:       http.MethodGet,
		ContentToSend: `[{
			"id": "{endpointId1}",
			"name": "endpoint1"
		}, {
			"id": "{endpointId2}",
			"name": "endpoint2"
		}]`}})
	defer server.Close()
	result, err := api.GetDomainEndpoints("123")
	if err != nil {
		t.Error("Failed call of GetDomainEndpoints()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetDomainEndpointsFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/domains/123/endpoints",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetDomainEndpoints("123") })
}

func TestCreateDomainEndpoint(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/domains/123/endpoints",
		Method:           http.MethodPost,
		EstimatedContent: `{"name":"endpoint"}`,
		HeadersToSend:    map[string]string{"Location": "/v1/users/{userId}/domain/123/endpoints/456"}}})
	defer server.Close()
	id, err := api.CreateDomainEndpoint("123", &DomainEndpointData{Name: "endpoint"})
	if err != nil {
		t.Error("Failed call of CreateDomainEndpoint()")
		return
	}
	expect(t, id, "456")
}

func TestCreateDomainEndpointFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:    "/v1/users/userId/domains/123/endpoints",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateDomainEndpoint("123", &DomainEndpointData{Name: "endpoint"})
	})
}

func TestUpdateDomainEndpoint(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/domains/123/endpoints/456",
		EstimatedContent: `{"name":"endpoint1"}`,
		Method:       http.MethodPost}})
	defer server.Close()
	err := api.UpdateDomainEndpoint("123", "456", &DomainEndpointData{Name: "endpoint1"})
	if err != nil {
		t.Error("Failed call of UpdateDomainEndpoint()")
		return
	}
}


func TestDeleteDomainEndpoint(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/domains/123/endpoints/456",
		Method:       http.MethodDelete}})
	defer server.Close()
	err := api.DeleteDomainEndpoint("123", "456")
	if err != nil {
		t.Error("Failed call of DeleteDomainEndpoint()")
		return
	}
}

func TestGetDomainEndpoint(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/domains/123/endpoints/456",
		Method:       http.MethodGet,
		ContentToSend: `{
			"id": "{endpointId2}",
			"name": "endpoint2"
		}`}})
	defer server.Close()
	result, err := api.GetDomainEndpoint("123", "456")
	if err != nil {
		t.Error("Failed call of GetDomainEndpoint()")
		return
	}
	expect(t, result.Name, "endpoint2")
}

func TestGetDomainEndpointFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/domains/123/endpoints/456",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetDomainEndpoint("123", "456") })
}

func TestCreateDomainEndpointToken(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/domains/123/endpoints/456/tokens",
		Method:       http.MethodPost,
		EstimatedContent: `null`,
		ContentToSend: `{"token": "123", "expires": 10}`}})
	defer server.Close()
	result, err := api.CreateDomainEndpointToken("123", "456")
	if err != nil {
		t.Error("Failed call of CreateDomainEndpointToken()")
		return
	}
	expect(t, result.Token, "123")
}

func TestCreateDomainEndpointTokenFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/domains/123/endpoints/456/tokens",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.CreateDomainEndpointToken("123", "456") })
}
