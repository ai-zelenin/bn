package spec

import "encoding/xml"

type BusinessRuleTask struct {
	XMLName           xml.Name                   `xml:"businessRuleTask"`
	Text              string                     `xml:",chardata"`
	ID                string                     `xml:"id,attr"`
	Name              string                     `xml:"name,attr"`
	Documentation     string                     `xml:"documentation"`
	ExtensionElements BusinessRuleTaskExtensions `xml:"extensionElements"`
	Incoming          string                     `xml:"incoming"`
	Outgoing          string                     `xml:"outgoing"`
}

type BusinessRuleTaskExtensions struct {
	XMLName        xml.Name       `xml:"extensionElements"`
	IoMapping      IoMapping      `xml:"ioMapping"`
	CalledDecision CalledDecision `xml:"calledDecision"`
}
