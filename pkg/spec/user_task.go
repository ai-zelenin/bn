package spec

import "encoding/xml"

type UserTask struct {
	XMLName xml.Name `xml:"userTask"`
	Task
	ExtensionElements *UserTaskExtensions `xml:"extensionElements"`
}

type UserTaskExtensions struct {
	XMLName              xml.Name             `xml:"extensionElements"`
	Text                 string               `xml:",chardata"`
	AssignmentDefinition AssignmentDefinition `xml:"assignmentDefinition"`
	FormDefinition       FormDefinition       `xml:"formDefinition"`
	IoMapping            IoMapping            `xml:"ioMapping"`
	TaskHeaders          TaskHeaders          `xml:"taskHeaders"`
}
