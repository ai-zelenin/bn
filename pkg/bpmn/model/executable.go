package model

import (
	"context"
	"fmt"
)

type Executable interface {
	Exec(ctx context.Context, startIDs []string, elementHandler ElementHandler)
	GetNodes(startIDs ...string) []*Node
}

type ExecutableProcess struct {
	StartID string
	ElementIndex[Element]
}

func NewExecutableProcess(start Element, index ElementIndex[Element]) *ExecutableProcess {
	return &ExecutableProcess{
		StartID:      start.ID(),
		ElementIndex: index,
	}
}

func (p *ExecutableProcess) GetNodes(startIDs ...string) ([]*Node, error) {
	startNodes := make([]*Node, 0, len(startIDs))
	for _, startID := range startIDs {
		element, ok := p.ElementIndex[startID]
		if !ok {
			return nil, fmt.Errorf("%w can not find ref %s", ErrInvalidBPMN, startID)

		}
		node := p.MakeNodeRecursive(element)
		startNodes = append(startNodes, node)
	}
	return startNodes, nil
}

func (p *ExecutableProcess) Exec(ctx context.Context, startIDs []string, elementHandler ElementHandler) {
	for _, startID := range startIDs {
		element, ok := p.ElementIndex[startID]
		if ok {
			node := p.MakeNodeRecursive(element)
			flow := NewFlow(NewID(), node)
			go flow.Run(ctx, elementHandler)
		}
	}
}

func (p *ExecutableProcess) HandleElement(ctx context.Context, el Element, handler ElementHandler) (isContinue bool) {
	if handler == nil {
		return
	}
	isContinue = true
	err := handler.HandleElement(ctx, el)
	if err != nil {
		isContinue = handler.HandleError(ctx, el.ID(), err)
	}
	return isContinue
}

func (p *ExecutableProcess) ExecInWorkerPool(ctx context.Context, startIDs []string, elementHandler ElementHandler) {
	if elementHandler == nil {
		return
	}
	type Step struct {
		ID   string
		Node *Node
	}
	var pl WorkerPayload[*Step] = func(ctx context.Context, wp *WorkerPool[*Step], val *Step) {
		el := val.Node.Element
		isContinue := p.HandleElement(ctx, el, elementHandler)
		if isContinue {
			for _, node := range val.Node.Outgoing {
				next := &Step{
					ID:   node.ID(),
					Node: node,
				}
				wp.Handle(next)
			}
		}
	}
	wp := NewWorkerPool[*Step](10, pl)
	go func() {
		for _, startID := range startIDs {
			element, ok := p.ElementIndex[startID]
			if !ok {
				err := fmt.Errorf("%w can not find ref %s", ErrInvalidBPMN, startID)
				if !elementHandler.HandleError(ctx, startID, err) {
					return
				}
			}
			node := p.MakeNodeRecursive(element)
			wp.Handle(&Step{
				ID:   startID,
				Node: node,
			})
		}
	}()
	wp.Start(ctx)
}

func (p *ExecutableProcess) Process(ctx context.Context, elementHandler ElementHandler, startIDs ...string) {
	if elementHandler == nil {
		return
	}
	for _, startID := range startIDs {
		element, ok := p.ElementIndex[startID]
		if !ok {
			err := fmt.Errorf("%w can not find ref %s", ErrInvalidBPMN, startID)
			if !elementHandler.HandleError(ctx, startID, err) {
				return
			}
		}
		go func(el Element) {
			err := elementHandler.HandleElement(ctx, el)
			if err != nil {
				if !elementHandler.HandleError(ctx, el.ID(), err) {
					return
				}
			}
			p.Process(ctx, elementHandler, el.OutRef()...)
		}(element)
	}
}
