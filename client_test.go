package bandwidth

import "testing"
import "fmt"

func TestNew(t *testing.T) {
	api, _ := New("userId", "apiToken", "apiSecret")
	if api.UserID != "userId" {
		t.Error("Missing UserId")
	}
	if api.APIToken != "apiToken" {
		t.Error("Missing APIToken")
	}
	if api.APISecret != "apiSecret" {
		t.Error("Missing APISecret")
	}
	if api.APIEndPoint != "https://api.catapult.inetwork.com" {
		t.Error("Invalid APIEndPoint")
	}
	if api.APIVersion != "v1" {
		t.Error("Invalid APIVersion")
	}
}

func TestNewWithVersion(t *testing.T) {
	api, _ := New("userId", "apiToken", "apiSecret", "v0")
	if api.UserID != "userId" {
		t.Error("Missing UserId")
	}
	if api.APIToken != "apiToken" {
		t.Error("Missing APIToken")
	}
	if api.APISecret != "apiSecret" {
		t.Error("Missing APISecret")
	}
	if api.APIEndPoint != "https://api.catapult.inetwork.com" {
		t.Error("Invalid APIEndPoint")
	}
	if api.APIVersion != "v0" {
		t.Error("Missing APIVersion")
	}
}

func TestNewWithEndpointAndVersion(t *testing.T) {
	api, _ := New("userId", "apiToken", "apiSecret", "v0", "endpoint")
	if api.UserID != "userId" {
		t.Error("Missing UserId")
	}
	if api.APIToken != "apiToken" {
		t.Error("Missing APIToken")
	}
	if api.APISecret != "apiSecret" {
		t.Error("Missing APISecret")
	}
	if api.APIEndPoint != "endpoint" {
		t.Error("Missing APIEndPoint")
	}
	if api.APIVersion != "v0" {
		t.Error("Missing APIVersion")
	}
}

func TestNewFail(t *testing.T) {
	_, err := New("", "apiToken", "apiSecret")
	if err == nil {
		t.Error("Should fail with missing UserId")
	}
	_, err = New("userId", "", "apiSecret")
	if err == nil {
		t.Error("Should fail with missing ApiToken")
	}
	_, err = New("userId", "apiToken", "")
	if err == nil {
		t.Error("Should fail with missing ApiSecret")
	}
}



func getAPI() *Client{
	api, _ := New("userId", "apiToken", "apiSecret")
	return api
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
	if req.URL.String() != "https://api.catapult.inetwork.com/v1/test" {
		t.Errorf("Invalid request url %s", req.URL.String())
	}
	if req.Method != "GET" {
		t.Errorf("Invalid request method %s", req.Method)
	}
	if req.Header.Get("Accept") != "application/json" {
		t.Errorf("Invalid header Accept %s", req.Header.Get("Accept"))
	}
	if req.Header.Get("User-Agent") != fmt.Sprintf("go-bandwidth-v%s", Version) {
		t.Errorf("Invalid header User-Agent %s", req.Header.Get("User-Agent"))
	}
	if req.Header.Get("Authorization") != "Basic YXBpVG9rZW46YXBpU2VjcmV0" {
		t.Errorf("Invalid header Authorization %s", req.Header.Get("Authorization"))
	}
}
