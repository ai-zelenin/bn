package spec

import "encoding/xml"

type ServiceTask struct {
	XMLName                          xml.Name                         `xml:"serviceTask"`
	Text                             string                           `xml:",chardata"`
	ID                               string                           `xml:"id,attr"`
	Name                             string                           `xml:"name,attr"`
	Documentation                    string                           `xml:"documentation"`
	ExtensionElements                ServiceTaskExtensions            `xml:"extensionElements"`
	Incoming                         string                           `xml:"incoming"`
	Outgoing                         []string                         `xml:"outgoing"`
	MultiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics"`
}

type ServiceTaskExtensions struct {
	XMLName        xml.Name       `xml:"extensionElements"`
	Text           string         `xml:",chardata"`
	TaskDefinition TaskDefinition `xml:"taskDefinition"`
	IoMapping      IoMapping      `xml:"ioMapping"`
	TaskHeaders    TaskHeaders    `xml:"taskHeaders"`
}
