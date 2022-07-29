package spec

import "encoding/xml"

type ManualTask struct {
	XMLName       xml.Name `xml:"manualTask"`
	Text          string   `xml:",chardata"`
	ID            string   `xml:"id,attr"`
	Name          string   `xml:"name,attr"`
	Documentation string   `xml:"documentation"`
	Incoming      string   `xml:"incoming"`
	Outgoing      string   `xml:"outgoing"`
}
