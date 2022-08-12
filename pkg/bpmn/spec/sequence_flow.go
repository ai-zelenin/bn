package spec

import "encoding/xml"

type SequenceFlow struct {
	XMLName             xml.Name            `xml:"sequenceFlow"`
	Text                string              `xml:",chardata"`
	Id                  string              `xml:"id,attr"`
	SourceRef           string              `xml:"sourceRef,attr"`
	TargetRef           string              `xml:"targetRef,attr"`
	Name                string              `xml:"name,attr"`
	ConditionExpression ConditionExpression `xml:"conditionExpression"`
}

func (t *SequenceFlow) ID() string {
	return t.Id
}

type ConditionExpression struct {
	XMLName xml.Name `xml:"conditionExpression"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
}
