package bandwidth

import (
	"net/http"
	"testing"
	"bytes"
	"io/ioutil"
)

func TestGetMediaFiles(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/media",
		Method:       http.MethodGet,
		ContentToSend: `[{
			"mediaName": "file1"
		}, {
			"mediaName": "file2"
		}]`}})
	defer server.Close()
	result, err := api.GetMediaFiles()
	if err != nil {
		t.Error("Failed call of GetMediaFiles()")
		return
	}
	expect(t, len(result), 2)
}

func TestGetMediaFilesFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/v1/users/userId/media",
		Method:           http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) { return api.GetMediaFiles() })
}


func TestDeleteMediaFile(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/media/file1",
		Method:       http.MethodDelete}})
	defer server.Close()
	err := api.DeleteMediaFile("file1")
	if err != nil {
		t.Error("Failed call of DeleteMediaFile()")
		return
	}
}

func TestUploadMediaFile(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/media/file1",
		Method:       http.MethodPut,
		EstimatedContent: "123",
		EstimatedHeaders: map[string]string {"Content-Type": "text/plain"}}})
	defer server.Close()
	err := api.UploadMediaFile("file1", nopCloser{bytes.NewReader([]byte("123"))}, "text/plain")
	if err != nil {
		t.Error("Failed call of UploadMediaFile()")
		return
	}
}

func TestUploadMediaFileWithDefaultContentType(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/media/file1",
		Method:       http.MethodPut,
		EstimatedContent: "123",
		EstimatedHeaders: map[string]string {"Content-Type": "application/octet-stream"}}})
	defer server.Close()
	err := api.UploadMediaFile("file1", nopCloser{bytes.NewReader([]byte("123"))})
	if err != nil {
		t.Error("Failed call of UploadMediaFile()")
		return
	}
}

func TestUploadMediaFileWithFilePath(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/media/file1",
		Method:       http.MethodPut,
		EstimatedContent: "1234",
		EstimatedHeaders: map[string]string {"Content-Type": "text/plain"}}})
	defer server.Close()
	err := api.UploadMediaFile("file1", "test.txt", "text/plain")
	if err != nil {
		t.Error("Failed call of UploadMediaFile()")
		return
	}
}

func TestDownloadMediaFile(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/media/file1",
		Method:       http.MethodGet,
		ContentToSend: "123",
		HeadersToSend: map[string]string {"Content-Type": "text/plain"}}})
	defer server.Close()
	reader, contentType, err := api.DownloadMediaFile("file1")
	if err != nil {
		t.Error("Failed call of DownloadMediaFile()")
		return
	}
	defer reader.Close()
	expect(t, contentType, "text/plain")
	data, _ := ioutil.ReadAll(reader)
	expect(t, string(data), "123\n")
}

func TestDownloadMediaFileFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery: "/v1/users/userId/media/file1",
		Method:       http.MethodGet,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	_, _, err := api.DownloadMediaFile("file1")
	if err == nil {
		t.Error("Should fail here")
	}
}
