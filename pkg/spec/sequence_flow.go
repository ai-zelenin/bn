package spec

import "encoding/xml"

type SequenceFlow struct {
	XMLName             xml.Name            `xml:"sequenceFlow"`
	Text                string              `xml:",chardata"`
	ID                  string              `xml:"id,attr"`
	SourceRef           string              `xml:"sourceRef,attr"`
	TargetRef           string              `xml:"targetRef,attr"`
	Name                string              `xml:"name,attr"`
	ConditionExpression ConditionExpression `xml:"conditionExpression"`
}

type ConditionExpression struct {
	XMLName xml.Name `xml:"conditionExpression"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
}
