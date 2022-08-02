package spec

import "encoding/xml"

type BusinessRuleTask struct {
	XMLName xml.Name `xml:"businessRuleTask"`
	Task
	ExtensionElements BusinessRuleTaskExtensions `xml:"extensionElements"`
}

type BusinessRuleTaskExtensions struct {
	XMLName        xml.Name       `xml:"extensionElements"`
	IoMapping      IoMapping      `xml:"ioMapping"`
	CalledDecision CalledDecision `xml:"calledDecision"`
}
