package bandwidth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	api, _ := New("userId", "apiToken", "apiSecret")
	expect(t, api.UserID, "userId")
	expect(t, api.APIToken, "apiToken")
	expect(t, api.APISecret, "apiSecret")
	expect(t, api.APIEndPoint, "https://api.catapult.inetwork.com")
	expect(t, api.APIVersion, "v1")
}

func TestNewWithVersion(t *testing.T) {
	api, _ := New("userId", "apiToken", "apiSecret", "v0")
	expect(t, api.UserID, "userId")
	expect(t, api.APIToken, "apiToken")
	expect(t, api.APISecret, "apiSecret")
	expect(t, api.APIEndPoint, "https://api.catapult.inetwork.com")
	expect(t, api.APIVersion, "v0")
}

func TestNewWithEndpointAndVersion(t *testing.T) {
	api, _ := New("userId", "apiToken", "apiSecret", "v0", "endpoint")
	expect(t, api.UserID, "userId")
	expect(t, api.APIToken, "apiToken")
	expect(t, api.APISecret, "apiSecret")
	expect(t, api.APIEndPoint, "endpoint")
	expect(t, api.APIVersion, "v0")
}

func TestNewFail(t *testing.T) {
	shouldFail(t, func() (interface{}, error) { return New("", "apiToken", "apiSecret") })
	shouldFail(t, func() (interface{}, error) { return New("userId", "", "apiSecret") })
	shouldFail(t, func() (interface{}, error) { return New("userID", "apiToken", "") })
}

func TestConcatUserPath(t *testing.T) {
	api := getAPI()
	if api.concatUserPath("test") != "/users/userId/test" {
		t.Error("Should return valid path (without slash)")
	}
	if api.concatUserPath("/test") != "/users/userId/test" {
		t.Error("Should return valid path (with slash)")
	}
}

func TestPrepareURL(t *testing.T) {
	api := getAPI()
	if api.prepareURL("test") != "https://api.catapult.inetwork.com/v1/test" {
		t.Error("Should return valid url (without slash)")
	}
	if api.prepareURL("/test") != "https://api.catapult.inetwork.com/v1/test" {
		t.Error("Should return valid url (with slash)")
	}
}

func TestCreateRequest(t *testing.T) {
	api := getAPI()
	req, err := api.createRequest(http.MethodGet, "/test")
	if err != nil {
		t.Fatal(err)
	}
	expect(t, req.URL.String(), "https://api.catapult.inetwork.com/v1/test")
	expect(t, req.Method, http.MethodGet)
	expect(t, req.Header.Get("Accept"), "application/json")
	expect(t, req.Header.Get("User-Agent"), fmt.Sprintf("go-bandwidth-v%s", Version))
	expect(t, req.Header.Get("Authorization"), "Basic YXBpVG9rZW46YXBpU2VjcmV0")
}

func TestCheckResponse(t *testing.T) {
	api := getAPI()
	data, _ := api.checkResponse(createFakeResponse(`{"test": "test"}`, 200))
	result := data.(map[string]interface{})
	expect(t, result["test"].(string), "test")
}

func TestCheckResponseFail(t *testing.T) {
	api := getAPI()
	err := shouldFail(t, func() (interface{}, error) {
		return api.checkResponse(createFakeResponse(`{"code": "400", "message": "some error"}`, 400))
	})
	expect(t, err.Error(), "some error")
	err = shouldFail(t, func() (interface{}, error) { return api.checkResponse(createFakeResponse(`{"code": "400"}`, 400)) })
	expect(t, err.Error(), "400")
	err = shouldFail(t, func() (interface{}, error) { return api.checkResponse(createFakeResponse("", 400)) })
	expect(t, err.Error(), "Http code 400")
}

func TestMakeRequest(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:  "/v1/test",
		ContentToSend: `{"test": "test"}`}})
	defer server.Close()
	result, _ := api.makeRequest(http.MethodGet, "/test")
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithQuery(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:  "/v1/test?field1=value1&field2=value+with+space",
		ContentToSend: `{"test": "test"}`}})
	defer server.Close()
	result, _ := api.makeRequest(http.MethodGet, "/test", map[string]interface{}{
		"field1": "value1",
		"field2": "value with space"})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithBody(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/test",
		Method:           http.MethodPost,
		EstimatedHeaders: map[string]string{"Content-Type": "application/json"},
		EstimatedContent: `{"field1":"value1","field2":"value with space"}`,
		ContentToSend:    `{"test": "test"}`}})
	defer server.Close()
	result, _ := api.makeRequest(http.MethodPost, "/test", map[string]interface{}{
		"field1": "value1",
		"field2": "value with space"})
	expect(t, result.(map[string]interface{})["test"], "test")
}
