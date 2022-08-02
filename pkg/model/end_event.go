package model

import "github.com/ai.zelenin/bpmn/pkg/spec"

type EndEvent struct {
	Spec *spec.EndEvent
}

func (e *EndEvent) ID() string {
	return e.Spec.Id
}

func (e *EndEvent) Outgoing() []string {
	return nil
}

func (e *EndEvent) Incoming() []string {
	return e.Spec.Incoming
}

func NewEndEvent(def *spec.EndEvent) (*EndEvent, error) {
	return &EndEvent{
		Spec: def,
	}, nil
}
