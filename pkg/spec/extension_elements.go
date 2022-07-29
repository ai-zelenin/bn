package spec

import "encoding/xml"

type AllExtensionElements struct {
	XMLName      xml.Name     `xml:"extensionElements"`
	Text         string       `xml:",chardata"`
	Subscription Subscription `xml:"subscription"`

	CalledDecision CalledDecision `xml:"calledDecision"`
	UserTaskForm   UserTaskForm   `xml:"userTaskForm"`

	AssignmentDefinition AssignmentDefinition `xml:"assignmentDefinition"`
	FormDefinition       FormDefinition       `xml:"formDefinition"`
	IoMapping            IoMapping            `xml:"ioMapping"`
	TaskHeaders          TaskHeaders          `xml:"taskHeaders"`
	TaskDefinition       TaskDefinition       `xml:"taskDefinition"`
}

type Subscription struct {
	XMLName        xml.Name `xml:"subscription"`
	Text           string   `xml:",chardata"`
	CorrelationKey string   `xml:"correlationKey,attr"`
}

type UserTaskForm struct {
	XMLName xml.Name `xml:"userTaskForm"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr"`
}

type IoMapping struct {
	XMLName xml.Name `xml:"ioMapping"`
	Text    string   `xml:",chardata"`
	Input   []struct {
		Text   string `xml:",chardata"`
		Source string `xml:"source,attr"`
		Target string `xml:"target,attr"`
	} `xml:"input"`
	Output []struct {
		Text   string `xml:",chardata"`
		Source string `xml:"source,attr"`
		Target string `xml:"target,attr"`
	} `xml:"output"`
}

type AssignmentDefinition struct {
	XMLName         xml.Name `xml:"assignmentDefinition"`
	Text            string   `xml:",chardata"`
	Assignee        string   `xml:"assignee,attr"`
	CandidateGroups string   `xml:"candidateGroups,attr"`
}

type FormDefinition struct {
	XMLName xml.Name `xml:"formDefinition"`
	Text    string   `xml:",chardata"`
	FormKey string   `xml:"formKey,attr"`
}

type TaskDefinition struct {
	XMLName xml.Name `xml:"taskDefinition"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
	Retries string   `xml:"retries,attr"`
}

type CalledDecision struct {
	XMLName        xml.Name `xml:"calledDecision"`
	Text           string   `xml:",chardata"`
	DecisionId     string   `xml:"decisionId,attr"`
	ResultVariable string   `xml:"resultVariable,attr"`
}
