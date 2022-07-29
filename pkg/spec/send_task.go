package spec

import "encoding/xml"

type SendTask struct {
	XMLName           xml.Name           `xml:"sendTask"`
	Text              string             `xml:",chardata"`
	ID                string             `xml:"id,attr"`
	Name              string             `xml:"name,attr"`
	Documentation     string             `xml:"documentation"`
	ExtensionElements SendTaskExtensions `xml:"extensionElements"`
	Incoming          string             `xml:"incoming"`
	Outgoing          string             `xml:"outgoing"`
}

type SendTaskExtensions struct {
	XMLName        xml.Name       `xml:"extensionElements"`
	TaskDefinition TaskDefinition `xml:"taskDefinition"`
	IoMapping      IoMapping      `xml:"ioMapping"`
	TaskHeaders    TaskHeaders    `xml:"taskHeaders"`
}
