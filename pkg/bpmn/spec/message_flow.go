package spec

import "encoding/xml"

type MessageFlow struct {
	XMLName       xml.Name `xml:"messageFlow"`
	Text          string   `xml:",chardata"`
	Id            string   `xml:"id,attr"`
	Name          string   `xml:"name,attr"`
	SourceRef     string   `xml:"sourceRef,attr"`
	TargetRef     string   `xml:"targetRef,attr"`
	Documentation string   `xml:"documentation"`
}

func (t *MessageFlow) ID() string {
	return t.Id
}
