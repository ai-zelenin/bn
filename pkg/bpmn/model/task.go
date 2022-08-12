package model

import (
	"fmt"
	spec2 "github.com/ai.zelenin/bpmn/pkg/bpmn/spec"
)

type Task struct {
	Spec                 *spec2.Task
	ServiceTaskSpec      *spec2.ServiceTask
	ScriptTaskSpec       *spec2.ScriptTask
	UserTaskSpec         *spec2.UserTask
	ManualTaskSpec       *spec2.ManualTask
	SendTaskSpec         *spec2.SendTask
	ReceiveTaskSpec      *spec2.ReceiveTask
	BusinessRuleTaskSpec *spec2.BusinessRuleTask
	CallActivitySpec     *spec2.CallActivity
}

func NewTask(definition spec2.Activity) (*Task, error) {
	task := &Task{
		Spec: definition.Definition(),
	}
	switch v := definition.(type) {
	case *spec2.ServiceTask:
		task.ServiceTaskSpec = v
	case *spec2.ScriptTask:
		task.ScriptTaskSpec = v
	case *spec2.UserTask:
		task.UserTaskSpec = v
	case *spec2.ManualTask:
		task.ManualTaskSpec = v
	case *spec2.SendTask:
		task.SendTaskSpec = v
	case *spec2.ReceiveTask:
		task.ReceiveTaskSpec = v
	case *spec2.BusinessRuleTask:
		task.BusinessRuleTaskSpec = v
	case *spec2.CallActivity:
		task.CallActivitySpec = v
	default:
		return nil, fmt.Errorf("%w unexpected task type %T", ErrInvalidBPMN, definition)
	}
	return task, nil
}

func AddTasks[T spec2.Activity](taskMap map[string]*Task, definitions ...T) error {
	for _, definition := range definitions {
		t, err := NewTask(definition)
		if err != nil {
			return err
		}
		taskMap[t.ID()] = t
	}
	return nil
}

func (t *Task) ID() string {
	return t.Spec.Id
}

func (t *Task) OutRef() []string {
	return t.Spec.Outgoing
}

func (t *Task) InRef() []string {
	return t.Spec.Incoming
}
