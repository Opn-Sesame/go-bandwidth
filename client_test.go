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
}

func TestNewWithEndpointAndVersion(t *testing.T) {
	api, _ := New("userId", "apiToken", "apiSecret", "endpoint")
	expect(t, api.UserID, "userId")
	expect(t, api.APIToken, "apiToken")
	expect(t, api.APISecret, "apiSecret")
	expect(t, api.APIEndPoint, "endpoint")
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
	if api.prepareURL("test", "v1") != "https://api.catapult.inetwork.com/v1/test" {
		t.Error("Should return valid url (without slash)")
	}
	if api.prepareURL("/test", "v1") != "https://api.catapult.inetwork.com/v1/test" {
		t.Error("Should return valid url (with slash)")
	}
}

func TestCreateRequest(t *testing.T) {
	api := getAPI()
	req, err := api.createRequest(http.MethodGet, "/test", "v1")
	if err != nil {
		t.Fatal(err)
	}
	expect(t, req.URL.String(), "https://api.catapult.inetwork.com/v1/test")
	expect(t, req.Method, http.MethodGet)
	expect(t, req.Header.Get("Accept"), "application/json")
	expect(t, req.Header.Get("User-Agent"), fmt.Sprintf("go-bandwidth/v%s", Version))
	expect(t, req.Header.Get("Authorization"), "Basic YXBpVG9rZW46YXBpU2VjcmV0")
}

func TestCreateRequestFail(t *testing.T) {
	api := getAPI()
	shouldFail(t, func() (interface{}, error) {
		return api.createRequest("invalid\n\r\tmethod", "invalid:/\n/ url = ", "v1")
	})
}

func TestCheckResponse(t *testing.T) {
	type Test struct {
		Test string `json:"test"`
	}
	api := getAPI()
	data, _, _ := api.checkResponse(createFakeResponse(`{"test": "test"}`, 200), map[string]interface{}{})
	result := data.(map[string]interface{})
	expect(t, result["test"].(string), "test")
	data, _, _ = api.checkResponse(createFakeResponse(`{"test": "test"}`, 200), nil)
	result = data.(map[string]interface{})
	expect(t, result["test"].(string), "test")
	data, _, _ = api.checkResponse(createFakeResponse(`{"test": "test"}`, 200), &Test{})
	testResult := data.(*Test)
	expect(t, testResult.Test, "test")
}

func TestCheckResponseFail(t *testing.T) {
	api := getAPI()
	fail := func(action func() (interface{}, http.Header, error)) error {
		_, _, err := action()
		if err == nil {
			t.Error("Should fail here")
		}
		return err
	}
	err := fail(func() (interface{}, http.Header, error) {
		return api.checkResponse(createFakeResponse(`{"code": "400", "message": "some error"}`, 400), nil)
	})
	expect(t, err.Error(), "some error")
	err = fail(func() (interface{}, http.Header, error) {
		return api.checkResponse(createFakeResponse(`{"code": "400"}`, 400), nil)
	})
	expect(t, err.Error(), "400")
	err = fail(func() (interface{}, http.Header, error) {
		return api.checkResponse(createFakeResponse("", 400), nil)
	})
	expect(t, err.Error(), "Http code 400")
	fail(func() (interface{}, http.Header, error) {
		return api.checkResponse(createFakeResponse("invalid\njson", 400), nil)
	})
}

func TestMakeRequest(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:  "/v1/test",
		ContentToSend: `{"test": "test"}`}})
	defer server.Close()
	result, _, _ := api.makeRequest(http.MethodGet, "/test", map[string]interface{}{})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestFail(t *testing.T) {
	api := getAPI()
	shouldFail(t, func() (interface{}, error) {
		_, _, err := api.makeRequest("invalid\n\r\tmethod", "invalid:/\n/ url = ")
		return nil, err
	})
}

func TestMakeRequestWithArrayAsResponse(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:  "/v1/test",
		ContentToSend: `[{"test": "test"}]`}})
	defer server.Close()
	result, _, _ := api.makeRequest(http.MethodGet, "/test", &[]map[string]string{})
	list := *(result.(*[]map[string]string))
	expect(t, len(list), 1)
	expect(t, list[0]["test"], "test")
}

func TestMakeRequestWithQuery(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:  "/v1/test?field1=value1&field2=value+with+space",
		ContentToSend: `{"test": "test"}`}})
	defer server.Close()
	result, _, _ := api.makeRequest(http.MethodGet, "/test", nil, map[string]string{
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
	result, _, _ := api.makeRequest(http.MethodPost, "/test", nil, map[string]interface{}{
		"field1": "value1",
		"field2": "value with space"})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithEmptyResponse(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:  "/v1/test",
		Method:        http.MethodGet,
		ContentToSend: ""}})
	defer server.Close()
	result, _, _ := api.makeRequest(http.MethodGet, "/test", &[]interface{}{})
	expect(t, len(*result.(*[]interface{})), 0)
}

func TestGetIDFromLocationHeader(t *testing.T) {
	headers := http.Header{"Location": []string{"http://localhost/123"}}
	headers = http.Header{"Location": []string{""}}
	expect(t, getIDFromLocationHeader(headers), "")
	headers = http.Header{}
	expect(t, getIDFromLocationHeader(headers), "")
}

func TestGetIDFromLocation(t *testing.T) {
	expect(t, getIDFromLocation("http://localhost/123"), "123")
	expect(t, getIDFromLocation(""), "")
}
