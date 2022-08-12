package model

import (
	"github.com/ai.zelenin/bpmn/pkg/bpmn/spec"
)

type StartEvent struct {
	Spec *spec.StartEvent
}

func NewStartEvent(def *spec.StartEvent) *StartEvent {
	return &StartEvent{
		Spec: def,
	}
}

func (s *StartEvent) ID() string {
	return s.Spec.ID()
}

func (s *StartEvent) OutRef() []string {
	return s.Spec.Outgoing
}

func (s *StartEvent) InRef() []string {
	return nil
}
