package spec

type Process struct {
	Text              string             `xml:",chardata"`
	Id                string             `xml:"id,attr"`
	Name              string             `xml:"name,attr"`
	IsExecutable      bool               `xml:"isExecutable,attr"`
	Documentation     string             `xml:"documentation"`
	ExtensionElements *ProcessExtensions `xml:"extensionElements"`

	StartEvent []*StartEvent `xml:"startEvent"`
	EndEvent   []*EndEvent   `xml:"endEvent"`

	ServiceTask      []*ServiceTask      `xml:"serviceTask"`
	ScriptTask       []*ScriptTask       `xml:"scriptTask"`
	UserTask         []*UserTask         `xml:"userTask"`
	CallActivity     []*CallActivity     `xml:"callActivity"`
	SendTask         []*SendTask         `xml:"sendTask"`
	ReceiveTask      []*ReceiveTask      `xml:"receiveTask"`
	ManualTask       []*ManualTask       `xml:"manualTask"`
	BusinessRuleTask []*BusinessRuleTask `xml:"businessRuleTask"`

	ExclusiveGateway []*ExclusiveGateway `xml:"exclusiveGateway"`
	ParallelGateway  []*ParallelGateway  `xml:"parallelGateway"`

	IntermediateCatchEvent []*IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	EventBasedGateway      []*EventBasedGateway      `xml:"eventBasedGateway"`
	IntermediateThrowEvent []*IntermediateThrowEvent `xml:"intermediateThrowEvent"`

	SubProcess    []*SubProcess    `xml:"subProcess"`
	BoundaryEvent []*BoundaryEvent `xml:"boundaryEvent"`
	SequenceFlow  []*SequenceFlow  `xml:"sequenceFlow"`
}

func (t *Process) ID() string {
	return t.Id
}

type ProcessExtensions struct {
	UserTaskForm UserTaskForm `xml:"userTaskForm"`
}
