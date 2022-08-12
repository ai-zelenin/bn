package spec

import "encoding/xml"

type EndEvent struct {
	XMLName           xml.Name           `xml:"endEvent"`
	Text              string             `xml:",chardata"`
	Id                string             `xml:"id,attr"`
	Name              string             `xml:"name,attr"`
	Documentation     string             `xml:"documentation"`
	ExtensionElements EndEventExtensions `xml:"extensionElements"`
	Incoming          []string           `xml:"incoming"`
}

type EndEventExtensions struct {
	XMLName   xml.Name  `xml:"extensionElements"`
	Text      string    `xml:",chardata"`
	IoMapping IoMapping `xml:"ioMapping"`
}
