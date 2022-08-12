package model

import (
	"github.com/ai.zelenin/bpmn/pkg/bpmn/spec"
)

type EndEvent struct {
	Spec *spec.EndEvent
}

func NewEndEvent(def *spec.EndEvent) *EndEvent {
	return &EndEvent{
		Spec: def,
	}
}

func (e *EndEvent) ID() string {
	return e.Spec.Id
}

func (e *EndEvent) OutRef() []string {
	return nil
}

func (e *EndEvent) InRef() []string {
	return e.Spec.Incoming
}
