package spec

import "encoding/xml"

type ExclusiveGateway struct {
	XMLName       xml.Name `xml:"exclusiveGateway"`
	Text          string   `xml:",chardata"`
	Id            string   `xml:"id,attr"`
	Name          string   `xml:"name,attr"`
	Documentation string   `xml:"documentation"`
	Incoming      string   `xml:"incoming"`
	Outgoing      []string `xml:"outgoing"`
}

func (t *ExclusiveGateway) ID() string {
	return t.Id
}
