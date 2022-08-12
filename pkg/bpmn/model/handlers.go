package model

import "context"

type (
	HandleElementFunc      func(ctx context.Context, el Element) (err error)
	HandleStartEventFunc   func(ctx context.Context, se *StartEvent) (err error)
	HandleSequenceFlowFunc func(ctx context.Context, sf *SequenceFlow) (err error)
	HandleTaskFunc         func(ctx context.Context, task *Task) (err error)
	HandleEndEventFunc     func(ctx context.Context, ee *EndEvent) (err error)
	HandleErrorFunc        func(ctx context.Context, elID string, err error) (isContinue bool)
)

type Hook interface {
	BeforeElement(ctx context.Context, el Element) (context.Context, error)
	AfterElement(ctx context.Context, el Element)
}

type ElementHandler interface {
	HandleElement(ctx context.Context, el Element) error
	HandleError(ctx context.Context, elID string, err error) (isContinue bool)
}
type StartEventHandler interface {
	HandleStartEvent(ctx context.Context, se *StartEvent) error
}
type SequenceFlowHandler interface {
	HandleSequenceFlow(ctx context.Context, sf *SequenceFlow) error
}
type TaskHandler interface {
	HandleTask(ctx context.Context, task *Task) error
}
type EndEventHandler interface {
	HandleEndEvent(ctx context.Context, ee *EndEvent) error
}
