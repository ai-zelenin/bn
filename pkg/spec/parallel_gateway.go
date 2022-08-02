package spec

import "encoding/xml"

type ParallelGateway struct {
	XMLName       xml.Name `xml:"parallelGateway"`
	Text          string   `xml:",chardata"`
	Id            string   `xml:"id,attr"`
	Name          string   `xml:"name,attr"`
	Documentation string   `xml:"documentation"`
	Incoming      []string `xml:"incoming"`
	Outgoing      []string `xml:"outgoing"`
}

func (t *ParallelGateway) ID() string {
	return t.Id
}
