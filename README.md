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

```golang
	api := bandwidth.New("userId", "apiToken", "apiSecret")
	callId, err := api.CallTo("+1-from-number", "+1-to-number")
```



See directory `samples` for more examples.

