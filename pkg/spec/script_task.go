package spec

import "encoding/xml"

type ScriptTask struct {
	XMLName           xml.Name             `xml:"scriptTask"`
	Text              string               `xml:",chardata"`
	ID                string               `xml:"id,attr"`
	Name              string               `xml:"name,attr"`
	Documentation     string               `xml:"documentation"`
	ExtensionElements ScriptTaskExtensions `xml:"extensionElements"`
	Incoming          string               `xml:"incoming"`
	Outgoing          string               `xml:"outgoing"`
}

type ScriptTaskExtensions struct {
	Text           string         `xml:",chardata"`
	TaskDefinition TaskDefinition `xml:"taskDefinition"`
	IoMapping      IoMapping      `xml:"ioMapping"`
	TaskHeaders    TaskHeaders    `xml:"taskHeaders"`
}
