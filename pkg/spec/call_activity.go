package spec

import "encoding/xml"

type CallActivity struct {
	XMLName           xml.Name               `xml:"callActivity"`
	Text              string                 `xml:",chardata"`
	ID                string                 `xml:"id,attr"`
	Name              string                 `xml:"name,attr"`
	Documentation     string                 `xml:"documentation"`
	ExtensionElements CallActivityExtensions `xml:"extensionElements"`
	Incoming          string                 `xml:"incoming"`
	Outgoing          string                 `xml:"outgoing"`
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
