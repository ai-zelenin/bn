package spec

import "encoding/xml"

type IntermediateCatchEvent struct {
	XMLName                xml.Name                         `xml:"intermediateCatchEvent"`
	Text                   string                           `xml:",chardata"`
	ID                     string                           `xml:"id,attr"`
	Name                   string                           `xml:"name,attr"`
	Documentation          string                           `xml:"documentation"`
	ExtensionElements      IntermediateCatchEventExtensions `xml:"extensionElements"`
	Incoming               string                           `xml:"incoming"`
	MessageEventDefinition MessageEventDefinition           `xml:"messageEventDefinition"`
	TimerEventDefinition   TimerEventDefinition             `xml:"timerEventDefinition"`
}

type IntermediateCatchEventExtensions struct {
	XMLName   xml.Name  `xml:"extensionElements"`
	Text      string    `xml:",chardata"`
	IoMapping IoMapping `xml:"ioMapping"`
}
