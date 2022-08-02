package spec

import "encoding/xml"

type ReceiveTask struct {
	XMLName xml.Name `xml:"receiveTask"`
	Task
	ExtensionElements ReceiveTaskExtensions `xml:"extensionElements"`
}

type ReceiveTaskExtensions struct {
	XMLName   xml.Name  `xml:"extensionElements"`
	IoMapping IoMapping `xml:"ioMapping"`
}
