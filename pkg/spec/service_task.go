package spec

import "encoding/xml"

type ServiceTask struct {
	XMLName xml.Name `xml:"serviceTask"`
	*Task
	ExtensionElements ServiceTaskExtensions `xml:"extensionElements"`
}

type ServiceTaskExtensions struct {
	XMLName        xml.Name       `xml:"extensionElements"`
	Text           string         `xml:",chardata"`
	TaskDefinition TaskDefinition `xml:"taskDefinition"`
	IoMapping      IoMapping      `xml:"ioMapping"`
	TaskHeaders    TaskHeaders    `xml:"taskHeaders"`
}
