package model

import (
	"context"
	"fmt"
	"github.com/ai.zelenin/bpmn/pkg/spec"
)

type Process struct {
	ID            string
	Spec          *spec.Process
	StartEvent    *StartEvent
	Tasks         map[string]*Task
	EndEvents     map[string]*EndEvent
	SequenceFlows map[string]*SequenceFlow
	Index         map[string]Element
}

func NewProcess(definition *spec.Process) (*Process, error) {
	if !definition.IsExecutable {
		return nil, fmt.Errorf("%w process %s is not executable", ErrInvalidBPMN, definition.ID())
	}
	if len(definition.StartEvent) != 1 {
		return nil, fmt.Errorf("%w wrong number of start events", ErrInvalidBPMN)
	}
	index := make(map[string]Element)
	startEvent, err := NewStartEvent(definition.StartEvent[0])
	if err != nil {
		return nil, err
	}
	index[startEvent.ID()] = startEvent

	endEvents := make(map[string]*EndEvent)
	for _, endEvent := range definition.EndEvent {
		ee, err := NewEndEvent(endEvent)
		if err != nil {
			return nil, err
		}
		endEvents[ee.ID()] = ee
		index[ee.ID()] = ee
	}
	tasks := make(map[string]*Task)
	err = AddTasks[*spec.ServiceTask](tasks, definition.ServiceTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec.ScriptTask](tasks, definition.ScriptTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec.UserTask](tasks, definition.UserTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec.CallActivity](tasks, definition.CallActivity...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec.SendTask](tasks, definition.SendTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec.ReceiveTask](tasks, definition.ReceiveTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec.ManualTask](tasks, definition.ManualTask...)
	if err != nil {
		return nil, err
	}
	err = AddTasks[*spec.BusinessRuleTask](tasks, definition.BusinessRuleTask...)
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		index[task.ID()] = task
	}
	sequenceFlows := make(map[string]*SequenceFlow)
	for _, seqFlow := range definition.SequenceFlow {
		sf, err := NewSequenceFlow(seqFlow)
		if err != nil {
			return nil, err
		}
		sequenceFlows[sf.ID()] = sf
		index[sf.ID()] = sf
	}
	return &Process{
		ID:            definition.ID(),
		Spec:          definition,
		StartEvent:    startEvent,
		EndEvents:     endEvents,
		SequenceFlows: sequenceFlows,
		Tasks:         tasks,
		Index:         index,
	}, nil
}

func (p *Process) AddHandler(id string) {

}

func AddTasks[T spec.Activity](taskMap map[string]*Task, definitions ...T) error {
	for _, definition := range definitions {
		t, err := NewTask(definition)
		if err != nil {
			return err
		}
		taskMap[t.ID()] = t
	}
	return nil
}

type ElementHandleFunc func(ctx context.Context, el Element) error

func (p *Process) Run(ctx context.Context, startFromID string, handleFunc ElementHandleFunc) error {
	if startFromID == "" {
		startFromID = p.StartEvent.ID()
	}
	el, ok := p.Index[startFromID]
	if !ok {
		return fmt.Errorf("%w can not find ref %s", ErrInvalidBPMN, startFromID)
	}
	if handleFunc != nil {
		err := handleFunc(ctx, el)
		if err != nil {
			return err
		}
	}
	outgoing := el.Outgoing()
	for _, out := range outgoing {
		err := p.Run(ctx, out, handleFunc)
		if err != nil {
			return err
		}
	}
	return nil
}
