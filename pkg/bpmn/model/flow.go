package model

import (
	"context"
)

type Flow struct {
	ID   string
	Node *Node
}

func NewFlow(ID string, node *Node) *Flow {
	return &Flow{
		ID:   ID,
		Node: node,
	}
}

func (f *Flow) Run(ctx context.Context, handler ElementHandler) {
	el := f.Node.Element
	var isContinue = true
	if handler != nil {
		err := handler.HandleElement(ctx, el)
		if err != nil {
			isContinue = handler.HandleError(ctx, el.ID(), err)
		}
	}
	if isContinue {
		out := f.Node.Out()
		switch len(out) {
		case 0:
			return
		case 1:
			f.Node = out[0]
			f.Run(ctx, handler)
		default:
			for _, node := range out {
				flow := f.Fork(NewID(), node)
				go flow.Run(ctx, handler)
			}
		}
	}
}

func (f *Flow) Fork(newFlowID string, node *Node) *Flow {
	return NewFlow(newFlowID, node)
}
