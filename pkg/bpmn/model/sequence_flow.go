package model

import (
	"github.com/ai.zelenin/bpmn/pkg/bpmn/spec"
)

type SequenceFlow struct {
	Node
	Spec *spec.SequenceFlow
}

func NewSequenceFlow(def *spec.SequenceFlow) *SequenceFlow {
	return &SequenceFlow{
		Spec: def,
	}
}

func (s *SequenceFlow) ID() string {
	return s.Spec.Id
}

func (s *SequenceFlow) OutRef() []string {
	return []string{s.Spec.TargetRef}
}

func (s *SequenceFlow) InRef() []string {
	return []string{s.Spec.SourceRef}
}
