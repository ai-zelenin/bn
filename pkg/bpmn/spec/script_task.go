package spec

import "encoding/xml"

type ScriptTask struct {
	XMLName xml.Name `xml:"scriptTask"`
	Task
	ExtensionElements ScriptTaskExtensions `xml:"extensionElements"`
}

type ScriptTaskExtensions struct {
	Text           string         `xml:",chardata"`
	TaskDefinition TaskDefinition `xml:"taskDefinition"`
	IoMapping      IoMapping      `xml:"ioMapping"`
	TaskHeaders    TaskHeaders    `xml:"taskHeaders"`
}
