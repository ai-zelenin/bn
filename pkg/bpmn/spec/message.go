package spec

import "encoding/xml"

type Message struct {
	XMLName           xml.Name          `xml:"message"`
	Text              string            `xml:",chardata"`
	Id                string            `xml:"id,attr"`
	Name              string            `xml:"name,attr"`
	ExtensionElements MessageExtensions `xml:"extensionElements"`
}

type MessageExtensions struct {
	XMLName      xml.Name     `xml:"extensionElements"`
	Text         string       `xml:",chardata"`
	Subscription Subscription `xml:"subscription"`
}

type MessageEventDefinition struct {
	XMLName    xml.Name `xml:"messageEventDefinition"`
	Text       string   `xml:",chardata"`
	Id         string   `xml:"id,attr"`
	MessageRef string   `xml:"messageRef,attr"`
}
