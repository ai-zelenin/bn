package spec

import "encoding/xml"

type CallActivity struct {
	XMLName xml.Name `xml:"callActivity"`
	Task
	ExtensionElements CallActivityExtensions `xml:"extensionElements"`
}

type CallActivityExtensions struct {
	XMLName       xml.Name      `xml:"extensionElements"`
	Text          string        `xml:",chardata"`
	CalledElement CalledElement `xml:"calledElement"`
	IoMapping     IoMapping     `xml:"ioMapping"`
}

type CalledElement struct {
	XMLName                    xml.Name `xml:"calledElement"`
	Text                       string   `xml:",chardata"`
	ProcessId                  string   `xml:"processId,attr"`
	PropagateAllChildVariables string   `xml:"propagateAllChildVariables,attr"`
}
