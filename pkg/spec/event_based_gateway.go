package spec

import "encoding/xml"

type EventBasedGateway struct {
	XMLName       xml.Name `xml:"eventBasedGateway"`
	Text          string   `xml:",chardata"`
	ID            string   `xml:"id,attr"`
	Documentation string   `xml:"documentation"`
	Incoming      []string `xml:"incoming"`
	Outgoing      []string `xml:"outgoing"`
}
