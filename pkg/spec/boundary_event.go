package spec

import "encoding/xml"

type BoundaryEvent struct {
	XMLName                xml.Name                `xml:"boundaryEvent"`
	Text                   string                  `xml:",chardata"`
	Id                     string                  `xml:"id,attr"`
	Name                   string                  `xml:"name,attr"`
	AttachedToRef          string                  `xml:"attachedToRef,attr"`
	CancelActivity         string                  `xml:"cancelActivity,attr"`
	Documentation          string                  `xml:"documentation"`
	ExtensionElements      BoundaryEventExtensions `xml:"extensionElements"`
	Outgoing               string                  `xml:"outgoing"`
	ErrorEventDefinition   ErrorEventDefinition    `xml:"errorEventDefinition"`
	MessageEventDefinition MessageEventDefinition  `xml:"messageEventDefinition"`
}

func (t *BoundaryEvent) ID() string {
	return t.Id
}

type BoundaryEventExtensions struct {
	XMLName   xml.Name  `xml:"extensionElements"`
	Text      string    `xml:",chardata"`
	IoMapping IoMapping `xml:"ioMapping"`
}
