package bandwidth

import (
	"fmt"
	"net/http"
	"net/textproto"
	"testing"
)

func TestNew(t *testing.T) {
	api, _ := New(testAccountID, "apiToken", "apiSecret", "userName", "password", nil, nil)
	expect(t, api.AccountID, testAccountID)
	expect(t, api.APIToken, "apiToken")
	expect(t, api.APISecret, "apiSecret")
	expect(t, api.UserName, "userName")
	expect(t, api.Password, "password")
	expect(t, api.AccountsEndpoint, "https://dashboard.bandwidth.com/api/accounts/"+testAccountID)
	expect(t, api.MessagingEndpoint, fmt.Sprintf("https://messaging.bandwidth.com/api/v2/users/%s/messages", testAccountID))
}

func TestNewFail(t *testing.T) {
	shouldFail(t, func() (interface{}, error) { return New("", "apiToken", "apiSecret", "username", "password", nil, nil) })
	shouldFail(t, func() (interface{}, error) { return New("userId", "", "apiSecret", "username", "password", nil, nil) })
	shouldFail(t, func() (interface{}, error) { return New("userID", "apiToken", "", "username", "password", nil, nil) })
	shouldFail(t, func() (interface{}, error) { return New("userID", "apiToken", "apiSecret", "", "password", nil, nil) })
	shouldFail(t, func() (interface{}, error) { return New("userID", "apiToken", "apiSecret", "username", "", nil, nil) })
}

func TestCreateRequest(t *testing.T) {
	endpoint := "https://localhost"
	api := getAPI(endpoint)
	req, err := api.createRequest(http.MethodGet, endpoint+"/v2/test", messagingRequest)
	if err != nil {
		t.Fatal(err)
	}
	expect(t, req.URL.String(), endpoint+"/v2/test")
	expect(t, req.Method, http.MethodGet)
	expect(t, req.Header.Get("Accept"), "application/json")
	expect(t, req.Header.Get("User-Agent"), "go-bandwidth/v2")
	expect(t, req.Header.Get("Authorization"), "Basic YXBpVG9rZW46YXBpU2VjcmV0")
}

func TestCreateRequestFail(t *testing.T) {
	endpoint := "https://localhost"
	api := getAPI(endpoint)
	shouldFail(t, func() (interface{}, error) {
		return api.createRequest("invalid\n\r\tmethod", "invalid:/\n/ url = ", messagingRequest)
	})
}

func TestCheckJSONResponse(t *testing.T) {
	type Test struct {
		Test string `json:"test"`
	}
	endpoint := "https://localhost"
	api := getAPI(endpoint)
	data, _, _ := api.checkJSONResponse(createFakeResponse(`{"test": "test"}`, 200), map[string]interface{}{})
	result := data.(map[string]interface{})
	expect(t, result["test"].(string), "test")
	data, _, _ = api.checkJSONResponse(createFakeResponse(`{"test": "test"}`, 200), nil)
	result = data.(map[string]interface{})
	expect(t, result["test"].(string), "test")
	data, _, _ = api.checkJSONResponse(createFakeResponse(`{"test": "test"}`, 200), &Test{})
	testResult := data.(*Test)
	expect(t, testResult.Test, "test")
}

func TestCheckResponseFail(t *testing.T) {
	endpoint := "https://localhost"
	api := getAPI(endpoint)
	fail := func(action func() (interface{}, http.Header, error)) error {
		_, _, err := action()
		if err == nil {
			t.Error("Should fail here")
		}
		return err
	}
	err := fail(func() (interface{}, http.Header, error) {
		return api.checkJSONResponse(createFakeResponse(`{"code": "400", "message": "some error"}`, 400), nil)
	})
	expect(t, err.Error(), "some error")
	err = fail(func() (interface{}, http.Header, error) {
		return api.checkJSONResponse(createFakeResponse(`{"code": "400"}`, 400), nil)
	})
	expect(t, err.Error(), "400")
	err = fail(func() (interface{}, http.Header, error) {
		return api.checkJSONResponse(createFakeResponse("", 400), nil)
	})
	expect(t, err.Error(), "Http code 400")
	fail(func() (interface{}, http.Header, error) {
		return api.checkJSONResponse(createFakeResponse("invalid\njson", 400), nil)
	})
	err = fail(func() (interface{}, http.Header, error) {
		resp := createFakeResponse("", 429)
		resp.Header = map[string][]string{textproto.CanonicalMIMEHeaderKey("X-RateLimit-Reset"): []string{"1479308598680"}}
		return api.checkJSONResponse(resp, nil)
	})
	e := err.(*RateLimitError)
	expect(t, e.Reset.Unix(), int64(1479308599))
}

func TestMakeRequest(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:  fmt.Sprintf("/api/v2/users/%s/messages", testAccountID),
		ContentToSend: `{"test": "test"}`}})
	defer server.Close()
	result, _, _ := api.makeMessagingRequest(http.MethodGet, api.MessagingEndpoint, map[string]interface{}{})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithQuery(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:  fmt.Sprintf("/api/v2/users/%s/messages?field1=value1&field2=value+with+space", testAccountID),
		ContentToSend: `{"test": "test"}`}})
	defer server.Close()
	result, _, _ := api.makeMessagingRequest(http.MethodGet, api.MessagingEndpoint, nil, map[string]string{
		"field1": "value1",
		"field2": "value with space"})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithBody(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     fmt.Sprintf("/api/v2/users/%s/messages", testAccountID),
		Method:           http.MethodPost,
		EstimatedHeaders: map[string]string{"Content-Type": "application/json"},
		EstimatedContent: `{"field1":"value1","field2":"value with space"}`,
		ContentToSend:    `{"test": "test"}`}})
	defer server.Close()
	result, _, _ := api.makeMessagingRequest(http.MethodPost, api.MessagingEndpoint, nil, map[string]interface{}{
		"field1": "value1",
		"field2": "value with space"})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithEmptyResponse(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:  fmt.Sprintf("/api/v2/users/%s/messages", testAccountID),
		Method:        http.MethodGet,
		ContentToSend: ""}})
	defer server.Close()
	result, _, _ := api.makeMessagingRequest(http.MethodGet, api.MessagingEndpoint, &[]interface{}{})
	expect(t, len(*result.(*[]interface{})), 0)
}
