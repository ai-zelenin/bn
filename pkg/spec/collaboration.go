package spec

import "encoding/xml"

type Collaboration struct {
	XMLName       xml.Name      `xml:"collaboration"`
	Text          string        `xml:",chardata"`
	Id            string        `xml:"id,attr"`
	Documentation string        `xml:"documentation"`
	Participant   []Participant `xml:"participant"`
	MessageFlow   []MessageFlow `xml:"messageFlow"`
}

func (t *Collaboration) ID() string {
	return t.Id
}
