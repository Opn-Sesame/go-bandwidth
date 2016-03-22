Catapult API in Go [![GoDoc](https://godoc.org/bandwidthcom/go-bandwidth?status.svg)](https://godoc.org/github.com/bandwidthcom/go-bandwidth) [![Build Status](https://travis-ci.org/bandwidthcom/go-bandwidth.svg)](https://travis-ci.org/bandwidthcom/go-bandwidth)
===============


Bandwidth [Bandwidth's App Platform](http://ap.bandwidth.com/?utm_medium=social&utm_source=github&utm_campaign=dtolb&utm_content=) Go SDK

With go-bandwidth  you have access to the entire set of API methods including:
* **Account** - get user's account data and transactions,
* **Application** - manage user's applications,
* **AvailableNumber** - search free local or toll-free phone numbers,
* **Bridge** - control bridges between calls,
* **Call** - get access to user's calls,
* **Conference** - manage user's conferences,
* **ConferenceMember** - make actions with conference members,
* **Domain** - get access to user's domains,
* **EntryPoint** - control of endpoints of domains,
* **Error** - list of errors,
* **Media** - list, upload and download files to Bandwidth API server,
* **Message** - send SMS/MMS, list messages,
* **NumberInfo** - receive CNUM info by phone number,
* **PhoneNumber** - get access to user's phone numbers,
* **Recording** - mamange user's recordings.

Also you can work with Bandwidth XML using special types. 

## Install

```
     go get github.com/bandwidthcom/go-bandwidth
```


## Getting Started

* Install `go-bandwidth`,
* **Get user id, api token and secret** - to use the Catapult API you need these data.  You can get them [here](https://catapult.inetwork.com/pages/catapult.jsf) on the tab "Account",
* **Set user id, api token and secret**

```golang
	import "github.com/bandwidthcom/go-bandwidth"
	
	api := bandwidth.New("userId", "apiToken", "apiSecret")
```

Read [Catapult Api documentation](http://ap.bandwidth.com/) for more details

## Examples

*All examples assume you have already setup your auth data!*

List all calls from special number

```go
  list, _ := api.GetCalls(map[string]string{"from": "+19195551212"})
```

List all received messages

```go
  list, _ := api.GetMessages(map[string]string{"state": "received"})
```

Send SMS

```go
  api.SendMessageTo("+19195551212", "+191955512142", "Test")
  // or
  api.CreateMessage(map[string]interface{}{"from": "+19195551212", "to":"+191955512142", "text":"Test"})
```


Send some SMSes

```go
  api.CreateMessage([]map[string]interface{}{
	  map[string]interface{}{"from": "+19195551212", "to":"+191955512142", "text":"Test1"}, 
	  map[string]interface{}{"from": "+19195551212", "to":"+191955512143", "text":"Test2"}})
```

Upload file

```go
  api.UploadMediaFile("avatar.png", "/local/path/to/file.png", "image/png")
```

Make a call

```go
  api.CallTo("+19195551212", "+191955512142")
  // or
  api.CreateCall(map[string]interface{}{"from": "+19195551212",  "to": "+191955512142"})
```

Reject incoming call

```go
  api.RejectIncomingCall(callId)
```

Create a gather
```go
  api.CreateGather(callId, map[string]interface{}{"maxDigits": 3, "interDigitTimeout": 5, "prompt": map[string]string{"sentence": "Please enter 3 digits"}})
```

Start a conference
```go
  api.CreateConference(map[string]interface{}{"from": "+19195551212"})
```

Add a member to the conference

```go
  api.CreateConferenceMember(conferenceId, map[string]interface{}{"callId": "id_of_call_to_add_to_this_conference", "joinTone": true, "leavingTone": true})
```


Connect 2 calls to a bridge

```go
  api.CreateBridge(map[string]interface{}{"callIds": []string{callId1, callId2}})
```

Search available local numbers to buy

```go
  list, _ := api.GetAvailableNumbers(bandwidth.AvailableNumberTypeLocal, map[string]string{"city": "Cary", "state": "NC", "quantity": "3"})
```
Get CNAM info for a number

```go
  info, _ := api.GetNumberInfo("+19195551212")
```

Buy a phone number

```go
  api.ReservePhoneNumber("+19195551212")
  // or
  api.CreatePhoneNumber(map[string]interface{}{"number": "+19195551212"})
```

List recordings

```go
  list, _ := api.GetRecordings()
```

Generate Bandwidth XML
```go
   import (
	   "github.com/bandwidthcom/go-bandwidth/xml"
	   "fmt"
   )
   
   response := &xml.Response{}
   speakSentence := xml.SpeakSentence{Sentence = "Transferring your call, please wait.", Voice = "paul", Gender = "male", Locale = "en_US"}
   transfer := xml.Transfer{
        TransferTo = "+13032218749",
        TransferCallerId = "private",
        SpeakSentence = &SpeakSentence{
            Sentence = "Inner speak sentence.",
            Voice = "paul",
            Gender = "male",
            Locale = "en_US"}}
    hangup := xml.Hangup{}

    append(response.Verbs, speakSentence)
	append(response.Verbs, transfer)
	append(response.Verbs, hangup)


    //as alternative we can pass list of verbs as
    //response = &xml.Response{Verbs = []interface{}{speakSentence, transfer, hangup}}

    fmt.Println(response.ToXML())
```

See directory `examples` for more demos.

# Bugs/Issues
Please open an issue in this repository and we'll handle it directly. If you have any questions please contact us at openapi@bandwidth.com.

