package model

import (
	"context"
	"github.com/ai.zelenin/bpmn/pkg/bpmn/spec"
)

type BPMN struct {
	Version   string
	Spec      *spec.Definitions
	Processes map[string]*Process
}

func NewBPMN(definitions *spec.Definitions) (*BPMN, error) {
	processDefinitions := make(map[string]*Process)
	for _, process := range definitions.Process {
		p, err := NewProcess(process)
		if err != nil {
			return nil, err
		}
		processDefinitions[process.ID()] = p
	}
	return &BPMN{
		Spec:      definitions,
		Processes: processDefinitions,
	}, nil
}

func (b *BPMN) LoadExecutableModel(ctx context.Context, processID string) (Executable, error) {
	p, ok := b.Processes[processID]
	if !ok {
		return nil, ErrProcessIDNotFound
	}
	return NewExecutableProcess(p.StartEvent, p.CreateIndex()), nil
}
