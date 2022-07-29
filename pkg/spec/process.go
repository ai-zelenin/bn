package spec

type Process struct {
	Text                   string                   `xml:",chardata"`
	ID                     string                   `xml:"id,attr"`
	Name                   string                   `xml:"name,attr"`
	IsExecutable           string                   `xml:"isExecutable,attr"`
	Documentation          string                   `xml:"documentation"`
	ExtensionElements      ProcessExtensions        `xml:"extensionElements"`
	StartEvent             StartEvent               `xml:"startEvent"`
	SendTask               []SendTask               `xml:"sendTask"`
	ReceiveTask            []ReceiveTask            `xml:"receiveTask"`
	UserTask               []UserTask               `xml:"userTask"`
	ManualTask             []ManualTask             `xml:"manualTask"`
	BusinessRuleTask       []BusinessRuleTask       `xml:"businessRuleTask"`
	ServiceTask            []ServiceTask            `xml:"serviceTask"`
	ScriptTask             []ScriptTask             `xml:"scriptTask"`
	CallActivity           []CallActivity           `xml:"callActivity"`
	ExclusiveGateway       []ExclusiveGateway       `xml:"exclusiveGateway"`
	ParallelGateway        []ParallelGateway        `xml:"parallelGateway"`
	IntermediateCatchEvent []IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	EventBasedGateway      []EventBasedGateway      `xml:"eventBasedGateway"`
	IntermediateThrowEvent []IntermediateThrowEvent `xml:"intermediateThrowEvent"`
	EndEvent               []EndEvent               `xml:"endEvent"`
	SubProcess             []SubProcess             `xml:"subProcess"`
	BoundaryEvent          []BoundaryEvent          `xml:"boundaryEvent"`
	SequenceFlow           []SequenceFlow           `xml:"sequenceFlow"`
}

type ProcessExtensions struct {
	UserTaskForm UserTaskForm `xml:"userTaskForm"`
}
