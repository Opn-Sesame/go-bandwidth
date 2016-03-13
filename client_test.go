package bandwidth

import "testing"

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
