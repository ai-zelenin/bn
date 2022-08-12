package spec

import "encoding/xml"

type TaskHeaders struct {
	XMLName xml.Name `xml:"taskHeaders"`
	Text    string   `xml:",chardata"`
	Header  []Header `xml:"header"`
}

type Header struct {
	XMLName xml.Name `xml:"header"`
	Text    string   `xml:",chardata"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
}
