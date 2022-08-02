package model

import (
	"fmt"
	"github.com/ai.zelenin/bpmn/pkg/spec"
)

type Task struct {
	Spec                 *spec.Task
	ServiceTaskSpec      *spec.ServiceTask
	ScriptTaskSpec       *spec.ScriptTask
	UserTaskSpec         *spec.UserTask
	ManualTaskSpec       *spec.ManualTask
	SendTaskSpec         *spec.SendTask
	ReceiveTaskSpec      *spec.ReceiveTask
	BusinessRuleTaskSpec *spec.BusinessRuleTask
	CallActivitySpec     *spec.CallActivity
}

func NewTask(definition spec.Activity) (*Task, error) {
	task := &Task{
		Spec: definition.Definition(),
	}
	switch v := definition.(type) {
	case *spec.ServiceTask:
		task.ServiceTaskSpec = v
	case *spec.ScriptTask:
		task.ScriptTaskSpec = v
	case *spec.UserTask:
		task.UserTaskSpec = v
	case *spec.ManualTask:
		task.ManualTaskSpec = v
	case *spec.SendTask:
		task.SendTaskSpec = v
	case *spec.ReceiveTask:
		task.ReceiveTaskSpec = v
	case *spec.BusinessRuleTask:
		task.BusinessRuleTaskSpec = v
	case *spec.CallActivity:
		task.CallActivitySpec = v
	default:
		return nil, fmt.Errorf("%w unexpected task type %T", ErrInvalidBPMN, definition)
	}
	return task, nil
}

func NewTaskFromAny(a any) (*Task, error) {
	te, ok := a.(spec.Activity)
	if !ok {
		return nil, fmt.Errorf("%w %T", ErrUnexpectedType, a)
	}
	return NewTask(te)
}

func (t *Task) ID() string {
	return t.Spec.Id
}

func (t *Task) Outgoing() []string {
	return t.Spec.Outgoing
}

func (t *Task) Incoming() []string {
	return t.Spec.Incoming
}
