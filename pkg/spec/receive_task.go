package spec

import "encoding/xml"

type ReceiveTask struct {
	XMLName           xml.Name              `xml:"receiveTask"`
	Text              string                `xml:",chardata"`
	ID                string                `xml:"id,attr"`
	Name              string                `xml:"name,attr"`
	MessageRef        string                `xml:"messageRef,attr"`
	Documentation     string                `xml:"documentation"`
	ExtensionElements ReceiveTaskExtensions `xml:"extensionElements"`
	Incoming          string                `xml:"incoming"`
	Outgoing          string                `xml:"outgoing"`
}

type ReceiveTaskExtensions struct {
	XMLName   xml.Name  `xml:"extensionElements"`
	IoMapping IoMapping `xml:"ioMapping"`
}
