package spec

import "encoding/xml"

type Participant struct {
	XMLName       xml.Name `xml:"participant"`
	Text          string   `xml:",chardata"`
	Id            string   `xml:"id,attr"`
	Name          string   `xml:"name,attr"`
	ProcessRef    string   `xml:"processRef,attr"`
	Documentation string   `xml:"documentation"`
}

func (t *Participant) ID() string {
	return t.Id
}
