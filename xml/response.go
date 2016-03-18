package xml

import (
    "encoding/xml"
	//"fmt"
)


// Response is response lemenet of BXML
type Response struct {
	Verbs []interface{} `xml:"."`
}

// ToXML build BXML as string
func (r *Response) ToXML() string{
	bytes, _ := xml.Marshal(r)
	return string(bytes)
}
