package model

import (
	"fmt"
	spec2 "github.com/ai.zelenin/bpmn/pkg/bpmn/spec"
)

type Process struct {
	Spec          *spec2.Process
	StartEvent    *StartEvent
	Tasks         map[string]*Task
	EndEvents     map[string]*EndEvent
	SequenceFlows map[string]*SequenceFlow
}

func NewProcess(definition *spec2.Process) (*Process, error) {
	if !definition.IsExecutable {
		return nil, fmt.Errorf("%w process %s is not executable", ErrInvalidBPMN, definition.ID())
	}
	if len(definition.StartEvent) != 1 {
		return nil, fmt.Errorf("%w wrong number of start events", ErrInvalidBPMN)
	}

	sequenceFlows := make(map[string]*SequenceFlow)
	for _, seqFlow := range definition.SequenceFlow {
		sf := NewSequenceFlow(seqFlow)
		sequenceFlows[sf.ID()] = sf
	}

	startEvent := NewStartEvent(definition.StartEvent[0])

	endEvents := make(map[string]*EndEvent)
	for _, endEvent := range definition.EndEvent {
		ee := NewEndEvent(endEvent)
		endEvents[ee.ID()] = ee
	}
	tasks := make(map[string]*Task)
	err := AddTasks[*spec2.ServiceTask](tasks, definition.ServiceTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec2.ScriptTask](tasks, definition.ScriptTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec2.UserTask](tasks, definition.UserTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec2.CallActivity](tasks, definition.CallActivity...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec2.SendTask](tasks, definition.SendTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec2.ReceiveTask](tasks, definition.ReceiveTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec2.ManualTask](tasks, definition.ManualTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec2.BusinessRuleTask](tasks, definition.BusinessRuleTask...)
	if err != nil {
		return nil, err
	}

	return &Process{
		Spec:          definition,
		StartEvent:    startEvent,
		EndEvents:     endEvents,
		SequenceFlows: sequenceFlows,
		Tasks:         tasks,
	}, nil
}

func (p *Process) CreateIndex() ElementIndex[Element] {
	globalMap := make(ElementIndex[Element])
	for _, sequenceFlow := range p.SequenceFlows {
		globalMap.AddElement(sequenceFlow)
	}
	globalMap.AddElement(p.StartEvent)
	for _, task := range p.Tasks {
		globalMap.AddElement(task)
	}
	for _, ee := range p.EndEvents {
		globalMap.AddElement(ee)
	}
	return globalMap
}

func (p *Process) ID() string {
	return p.Spec.ID()
}

func (p *Process) OutRef() []string {
	return []string{p.StartEvent.ID()}
}

func (p *Process) InRef() []string {
	return nil
}
