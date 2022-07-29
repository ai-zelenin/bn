package spec

import "encoding/xml"

type MultiInstanceLoopCharacteristics struct {
	XMLName             xml.Name                                  `xml:"multiInstanceLoopCharacteristics"`
	Text                string                                    `xml:",chardata"`
	IsSequential        string                                    `xml:"isSequential,attr"`
	ExtensionElements   MultiInstanceLoopCharacteristicsExtension `xml:"extensionElements"`
	CompletionCondition CompletionCondition                       `xml:"completionCondition"`
}

type MultiInstanceLoopCharacteristicsExtension struct {
	XMLName             xml.Name            `xml:"extensionElements"`
	Text                string              `xml:",chardata"`
	LoopCharacteristics LoopCharacteristics `xml:"loopCharacteristics"`
}

type LoopCharacteristics struct {
	XMLName          xml.Name `xml:"loopCharacteristics"`
	Text             string   `xml:",chardata"`
	InputCollection  string   `xml:"inputCollection,attr"`
	InputElement     string   `xml:"inputElement,attr"`
	OutputCollection string   `xml:"outputCollection,attr"`
	OutputElement    string   `xml:"outputElement,attr"`
}

type CompletionCondition struct {
	XMLName xml.Name `xml:"completionCondition"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
}
