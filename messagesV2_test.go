package bandwidth

import (
	"net/http"
	"testing"
)

func TestCreateMessageV2(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/api/v2/users/userId/messages",
		Method:           http.MethodPost,
		EstimatedContent: `{"from":"fromNumber","to":"toNumber","text":"text"}`,
		ContentToSend: `{
			"id"            : "14762070468292kw2fuqty55yp2b2",
			"time"          : "2016-09-14T18:20:16Z",
			"to"            : [
			  "+12345678902",
			  "+12345678903"
			],
			"from"          : "+12345678901",
			"text"          : "Hey, check this out!",
			"applicationId" : "93de2206-9669-4e07-948d-329f4b722ee2",
			"tag"           : "test message",
			"owner"         : "+12345678901",
			"media"         : [
			  "https://s3.amazonaws.com/bw-v2-api/demo.jpg"
			],
			"direction"     : "out",
			"segmentCount"  : 1
		  }`}})
	defer server.Close()
	message, err := api.CreateMessageV2(&CreateMessageDataV2{From: "fromNumber", To: "toNumber", Text: "text"})
	if err != nil {
		t.Error("Failed call of CreateMessage()")
		return
	}
	tm := message.Time.String()
	expect(t, message.ID, "14762070468292kw2fuqty55yp2b2")
	expect(t, tm, "2016-09-14 18:20:16 +0000 UTC")
	expect(t, len(message.To.([]interface{})), 2)
	expect(t, len(message.Media), 1)
}

func TestCreateMessageV2Fail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{RequestHandler{
		PathAndQuery:     "/api/v2/users/userId/messages",
		Method:           http.MethodPost,
		StatusCodeToSend: http.StatusBadRequest}})
	defer server.Close()
	shouldFail(t, func() (interface{}, error) {
		return api.CreateMessageV2(&CreateMessageDataV2{From: "fromNumber", To: "toNumber", Text: "text"})
	})
}
