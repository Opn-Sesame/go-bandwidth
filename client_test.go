package bandwidth

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func expect(t *testing.T, value interface{}, expected interface{}) {
	if value != expected {
		t.Errorf("Expected %v - Got %v", expected, value)
	}
}

func shouldFail(t *testing.T, action func() (interface{}, error)) error {
	_, err := action()
	if err == nil {
		t.Fatal("Should fail here")
		return nil
	}
	return err
}

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

func getAPI() *Client {
	api, _ := New("userId", "apiToken", "apiSecret")
	return api
}

func createFakeResponse(body string, statusCode int) *http.Response {
	return &http.Response{StatusCode: statusCode,
		Body: nopCloser{bytes.NewReader([]byte(body))}}
}

func startFakeServer(handler func(w http.ResponseWriter, r *http.Request)) (*httptest.Server, *Client) {
	api := getAPI()
	fakeServer := httptest.NewServer(http.HandlerFunc(handler))
	api.APIEndPoint = fakeServer.URL
	return fakeServer, api
}

func readText(t *testing.T, r io.Reader) string {
	text, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error("Error on reading content")
		return ""
	}
	return string(text)
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
	req, err := api.createRequest("GET", "/test")
	if err != nil {
		t.Fatal(err)
	}
	expect(t, req.URL.String(), "https://api.catapult.inetwork.com/v1/test")
	expect(t, req.Method, "GET")
	expect(t, req.Header.Get("Accept"), "application/json")
	expect(t, req.Header.Get("User-Agent"), fmt.Sprintf("go-bandwidth-v%s", Version))
	expect(t, req.Header.Get("Authorization"), "Basic YXBpVG9rZW46YXBpU2VjcmV0")
}

func TestCreateRequestFail(t *testing.T) {
	api := getAPI()
	shouldFail(t, func() (interface{}, error) { return api.createRequest("\r\nINVALID\nMETHOD", "/test") })
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
	fakeServer, api := startFakeServer(func(w http.ResponseWriter, r *http.Request) {
		expect(t, r.URL.String(), "/v1/test")
		expect(t, r.Method, "GET")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"test": "test"}`)
	})
	defer fakeServer.Close()
	result, _ := api.makeRequest("GET", "/test")
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithQuery(t *testing.T) {
	fakeServer, api := startFakeServer(func(w http.ResponseWriter, r *http.Request) {
		expect(t, r.URL.String(), "/v1/test?field1=value1&field2=value+with+space")
		expect(t, r.Method, "GET")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"test": "test"}`)
	})
	defer fakeServer.Close()
	result, _ := api.makeRequest("GET", "/test", map[string]interface{}{
		"field1": "value1",
		"field2": "value with space"})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithBody(t *testing.T) {
	fakeServer, api := startFakeServer(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		expect(t, r.URL.String(), "/v1/test")
		expect(t, r.Method, "POST")
		expect(t, r.Header.Get("Content-Type"), "application/json")
		expect(t, readText(t, r.Body), `{"field1":"value1","field2":"value with space"}`)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"test": "test"}`)
	})
	defer fakeServer.Close()
	result, _ := api.makeRequest("POST", "/test", map[string]interface{}{
		"field1": "value1",
		"field2": "value with space"})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestFail(t *testing.T) {
	api := getAPI()
	shouldFail(t, func() (interface{}, error) { return api.makeRequest("INVALID\nMETHOD\n", "/test") })
}
