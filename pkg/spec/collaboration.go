package spec

import "encoding/xml"

type Collaboration struct {
	XMLName       xml.Name      `xml:"collaboration"`
	Text          string        `xml:",chardata"`
	ID            string        `xml:"id,attr"`
	Documentation string        `xml:"documentation"`
	Participant   []Participant `xml:"participant"`
	MessageFlow   []MessageFlow `xml:"messageFlow"`
}
