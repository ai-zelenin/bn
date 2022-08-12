package spec

import "encoding/xml"

type SendTask struct {
	XMLName xml.Name `xml:"sendTask"`
	Task
	ExtensionElements SendTaskExtensions `xml:"extensionElements"`
}

type SendTaskExtensions struct {
	XMLName        xml.Name       `xml:"extensionElements"`
	TaskDefinition TaskDefinition `xml:"taskDefinition"`
	IoMapping      IoMapping      `xml:"ioMapping"`
	TaskHeaders    TaskHeaders    `xml:"taskHeaders"`
}
