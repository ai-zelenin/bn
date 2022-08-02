package model

import "github.com/ai.zelenin/bpmn/pkg/spec"

type SequenceFlow struct {
	Spec *spec.SequenceFlow
}

func NewSequenceFlow(def *spec.SequenceFlow) (*SequenceFlow, error) {
	return &SequenceFlow{
		Spec: def,
	}, nil
}

func (s *SequenceFlow) ID() string {
	return s.Spec.Id
}

func (s *SequenceFlow) Outgoing() []string {
	return []string{s.Spec.TargetRef}
}

func (s *SequenceFlow) Incoming() []string {
	return []string{s.Spec.SourceRef}
}
