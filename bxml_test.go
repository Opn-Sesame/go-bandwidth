package bandwidth

import (
	"encoding/xml"
	"testing"

	bxml "github.com/bandwidthcom/go-bandwidth/xml"
)

func TestToXML(t *testing.T) {
	response := &bxml.Response{}
	expect(t, response.ToXML(), `<Response></Response>`)
}

func TestToXMLWithSimpleVerb(t *testing.T) {
	type Test struct {
		XMLName xml.Name `xml:"Test"`
	}
	response := &bxml.Response{Verbs: []interface{}{Test{}}}
	expect(t, response.ToXML(), `<Response><Test></Test></Response>`)
}

func TestToXMLWithMultipleVerbs(t *testing.T) {
	type Test1 struct {
		XMLName xml.Name `xml:"Test1"`
	}
	type Test2 struct {
		XMLName xml.Name `xml:"Test2"`
	}
	response := &bxml.Response{Verbs: []interface{}{Test1{}, Test2{}}}
	expect(t, response.ToXML(), `<Response><Test1></Test1><Test2></Test2></Response>`)
}

func TestToXMLWithDefaultValues(t *testing.T) {
	type Test struct {
		XMLName xml.Name    `xml:"Test"`
		Field1  string      `xml:"field1,attr,omitempty"`
		Field2  interface{} `xml:"field2,attr,omitempty"`
		Field3  interface{} `xml:"field3,attr,omitempty"`
	}
	response := &bxml.Response{Verbs: []interface{}{Test{}}}
	expect(t, response.ToXML(), `<Response><Test></Test></Response>`)
	response = &bxml.Response{Verbs: []interface{}{Test{Field1: "value1", Field2: 11, Field3: false}}}
	expect(t, response.ToXML(), `<Response><Test field1="value1" field2="11" field3="false"></Test></Response>`)
	response = &bxml.Response{Verbs: []interface{}{Test{Field1: "value2", Field2: 12, Field3: true}}}
	expect(t, response.ToXML(), `<Response><Test field1="value2" field2="12" field3="true"></Test></Response>`)
}

func TestVerbs(t *testing.T) {
	response := &bxml.Response{Verbs: []interface{}{
		bxml.Gather{RequestURL: "url"},
		bxml.Pause{Duration: 10},
		bxml.Hangup{},
		bxml.PlayAudio{URL: "url"},
		bxml.Record{RequestURL: "url"},
		bxml.Redirect{RequestURL: "url"},
		bxml.Reject{Reason: "none"},
		bxml.SendMessage{From: "from", To: "to", Text: "text"},
		bxml.SpeakSentence{Sentence: "Hello"},
		bxml.Transfer{TransferTo: "number", SpeakSentence: &bxml.SpeakSentence{Sentence: "Please wait"}},
	}}
	expect(t, response.ToXML(), `<Response><Gather requestUrl="url"></Gather><Pause duration="10"></Pause><Hangup></Hangup><PlayAudio>url</PlayAudio><Record requestUrl="url"></Record><Redirect requestUrl="url"></Redirect><Reject reason="none"></Reject><SendMessage from="from" to="to">text</SendMessage><SpeakSentence>Hello</SpeakSentence><Transfer transferTo="number"><SpeakSentence>Please wait</SpeakSentence></Transfer></Response>`)
}
