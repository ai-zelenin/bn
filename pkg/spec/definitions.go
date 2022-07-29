package spec

import "encoding/xml"

type Definitions struct {
	XMLName                  xml.Name      `xml:"definitions"`
	Text                     string        `xml:",chardata"`
	ID                       string        `xml:"id,attr"`
	TargetNamespace          string        `xml:"targetNamespace,attr"`
	Exporter                 string        `xml:"exporter,attr"`
	ExporterVersion          string        `xml:"exporterVersion,attr"`
	ExecutionPlatform        string        `xml:"executionPlatform,attr"`
	ExecutionPlatformVersion string        `xml:"executionPlatformVersion,attr"`
	Message                  []Message     `xml:"message"`
	Error                    []Error       `xml:"error"`
	Collaboration            Collaboration `xml:"collaboration"`
	Process                  []Process     `xml:"process"`
	BPMNDiagram              []BPMNDiagram `xml:"BPMNDiagram"`
}
