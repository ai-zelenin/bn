package spec

import "encoding/xml"

type IntermediateThrowEvent struct {
	XMLName                xml.Name                         `xml:"intermediateThrowEvent"`
	Text                   string                           `xml:",chardata"`
	ID                     string                           `xml:"id,attr"`
	Name                   string                           `xml:"name,attr"`
	Documentation          string                           `xml:"documentation"`
	ExtensionElements      IntermediateThrowEventExtensions `xml:"extensionElements"`
	Incoming               string                           `xml:"incoming"`
	MessageEventDefinition MessageEventDefinition           `xml:"messageEventDefinition"`
}

type IntermediateThrowEventExtensions struct {
	XMLName        xml.Name       `xml:"extensionElements"`
	Text           string         `xml:",chardata"`
	IoMapping      IoMapping      `xml:"ioMapping"`
	TaskDefinition TaskDefinition `xml:"taskDefinition"`
	TaskHeaders    TaskHeaders    `xml:"taskHeaders"`
}
