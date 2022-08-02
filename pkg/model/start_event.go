package model

import (
	"github.com/ai.zelenin/bpmn/pkg/spec"
)

type StartEvent struct {
	Spec *spec.StartEvent
}

func NewStartEvent(def *spec.StartEvent) (*StartEvent, error) {
	return &StartEvent{
		Spec: def,
	}, nil
}

func (s *StartEvent) ID() string {
	return s.Spec.ID()
}

func (s *StartEvent) Outgoing() []string {
	return s.Spec.Outgoing
}

func (s *StartEvent) Incoming() []string {
	return nil
}
