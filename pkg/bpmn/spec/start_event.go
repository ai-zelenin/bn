package spec

import "encoding/xml"

type StartEvent struct {
	XMLName           xml.Name             `xml:"startEvent"`
	Text              string               `xml:",chardata"`
	Id                string               `xml:"id,attr"`
	Name              string               `xml:"name,attr"`
	Outgoing          []string             `xml:"outgoing"`
	Documentation     string               `xml:"documentation"`
	ExtensionElements StartEventExtensions `xml:"extensionElements"`
}

func (t *StartEvent) ID() string {
	return t.Id
}

type StartEventExtensions struct {
	XMLName   xml.Name  `xml:"extensionElements"`
	IoMapping IoMapping `xml:"ioMapping"`
}
