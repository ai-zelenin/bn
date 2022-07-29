package spec

import "encoding/xml"

type UserTask struct {
	XMLName           xml.Name           `xml:"userTask"`
	Text              string             `xml:",chardata"`
	ID                string             `xml:"id,attr"`
	Name              string             `xml:"name,attr"`
	Documentation     string             `xml:"documentation"`
	ExtensionElements UserTaskExtensions `xml:"extensionElements"`
	Incoming          string             `xml:"incoming"`
	Outgoing          string             `xml:"outgoing"`
}

type UserTaskExtensions struct {
	XMLName              xml.Name             `xml:"extensionElements"`
	Text                 string               `xml:",chardata"`
	AssignmentDefinition AssignmentDefinition `xml:"assignmentDefinition"`
	FormDefinition       FormDefinition       `xml:"formDefinition"`
	IoMapping            IoMapping            `xml:"ioMapping"`
	TaskHeaders          TaskHeaders          `xml:"taskHeaders"`
}
