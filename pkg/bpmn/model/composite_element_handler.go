package model

import (
	"context"
)

type CompositeElementHandler struct {
	GlobalHandleElementFuncs      []HandleElementFunc
	GlobalHandleStartEventFuncs   []HandleStartEventFunc
	GlobalHandleSequenceFlowFuncs []HandleSequenceFlowFunc
	GlobalHandleTaskFuncs         []HandleTaskFunc
	GlobalHandleEndEventFuncs     []HandleEndEventFunc
	HandleErrorFunc               HandleErrorFunc

	HandleElementFuncs      map[string][]HandleElementFunc
	HandleStartEventFuncs   map[string][]HandleStartEventFunc
	HandleSequenceFlowFuncs map[string][]HandleSequenceFlowFunc
	HandleTaskFuncs         map[string][]HandleTaskFunc
	HandleEndEventFuncs     map[string][]HandleEndEventFunc

	Hooks []Hook
}

func NewCompositeElementHandler() *CompositeElementHandler {
	return &CompositeElementHandler{
		GlobalHandleElementFuncs:      make([]HandleElementFunc, 0),
		GlobalHandleStartEventFuncs:   make([]HandleStartEventFunc, 0),
		GlobalHandleSequenceFlowFuncs: make([]HandleSequenceFlowFunc, 0),
		GlobalHandleTaskFuncs:         make([]HandleTaskFunc, 0),
		GlobalHandleEndEventFuncs:     make([]HandleEndEventFunc, 0),

		HandleElementFuncs:      make(map[string][]HandleElementFunc),
		HandleStartEventFuncs:   make(map[string][]HandleStartEventFunc, 0),
		HandleSequenceFlowFuncs: make(map[string][]HandleSequenceFlowFunc, 0),
		HandleTaskFuncs:         make(map[string][]HandleTaskFunc, 0),
		HandleEndEventFuncs:     make(map[string][]HandleEndEventFunc, 0),

		Hooks: make([]Hook, 0),
	}
}

func (c *CompositeElementHandler) HandleElement(ctx context.Context, el Element) (err error) {
	for _, cb := range c.GlobalHandleElementFuncs {
		if cb != nil {
			err = cb(ctx, el)
			if err != nil {
				return err
			}
		}
	}
	for _, hook := range c.Hooks {
		if hook != nil {
			ctx, err = hook.BeforeElement(ctx, el)
			if err != nil {
				return err
			}
		}
	}
	defer func() {
		for _, hook := range c.Hooks {
			if hook != nil {
				hook.AfterElement(ctx, el)
			}
		}
	}()
	for _, cb := range c.HandleElementFuncs[el.ID()] {
		if cb != nil {
			err = cb(ctx, el)
			if err != nil {
				return err
			}
		}
	}
	switch tEl := el.(type) {
	case *StartEvent:
		return c.HandleStartEvent(ctx, tEl)
	case *SequenceFlow:
		return c.HandleSequenceFlow(ctx, tEl)
	case *Task:
		return c.HandleTask(ctx, tEl)
	case *EndEvent:
		return c.HandleEndEvent(ctx, tEl)
	}
	return nil
}

func (c *CompositeElementHandler) HandleError(ctx context.Context, elID string, err error) (isContinue bool) {
	if c.HandleErrorFunc != nil {
		return c.HandleErrorFunc(ctx, elID, err)
	}
	return true
}

func (c *CompositeElementHandler) SetHandleError(cb HandleErrorFunc) {
	c.HandleErrorFunc = cb
}

func (c *CompositeElementHandler) HandleStartEvent(ctx context.Context, se *StartEvent) (err error) {
	for _, cb := range c.GlobalHandleStartEventFuncs {
		if cb != nil {
			err = cb(ctx, se)
			if err != nil {
				return err
			}
		}
	}
	for _, cb := range c.HandleStartEventFuncs[se.ID()] {
		if cb != nil {
			err = cb(ctx, se)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *CompositeElementHandler) HandleSequenceFlow(ctx context.Context, sf *SequenceFlow) (err error) {
	for _, cb := range c.GlobalHandleSequenceFlowFuncs {
		if cb != nil {
			err = cb(ctx, sf)
			if err != nil {
				return err
			}
		}
	}
	for _, cb := range c.HandleSequenceFlowFuncs[sf.ID()] {
		if cb != nil {
			err = cb(ctx, sf)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *CompositeElementHandler) HandleTask(ctx context.Context, task *Task) (err error) {
	for _, cb := range c.GlobalHandleTaskFuncs {
		if cb != nil {
			err = cb(ctx, task)
			if err != nil {
				return err
			}
		}
	}
	for _, cb := range c.HandleTaskFuncs[task.ID()] {
		if cb != nil {
			err = cb(ctx, task)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *CompositeElementHandler) HandleEndEvent(ctx context.Context, ee *EndEvent) (err error) {
	for _, cb := range c.GlobalHandleEndEventFuncs {
		if cb != nil {
			err = cb(ctx, ee)
			if err != nil {
				return err
			}
		}
	}
	for _, cb := range c.HandleEndEventFuncs[ee.ID()] {
		if cb != nil {
			err = cb(ctx, ee)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *CompositeElementHandler) AddHandleElementFunc(id string, cb HandleElementFunc) {
	funcs := c.HandleElementFuncs[id]
	funcs = append(funcs, cb)
	c.HandleElementFuncs[id] = funcs
}

func (c *CompositeElementHandler) AddHandleStartEventFunc(id string, cb HandleStartEventFunc) {
	funcs := c.HandleStartEventFuncs[id]
	funcs = append(funcs, cb)
	c.HandleStartEventFuncs[id] = funcs
}

func (c *CompositeElementHandler) AddHandleSequenceFlowFunc(id string, cb HandleSequenceFlowFunc) {
	funcs := c.HandleSequenceFlowFuncs[id]
	funcs = append(funcs, cb)
	c.HandleSequenceFlowFuncs[id] = funcs
}

func (c *CompositeElementHandler) AddHandleTaskFunc(id string, cb HandleTaskFunc) {
	funcs := c.HandleTaskFuncs[id]
	funcs = append(funcs, cb)
	c.HandleTaskFuncs[id] = funcs
}

func (c *CompositeElementHandler) AddHandleEndEventFunc(id string, cb HandleEndEventFunc) {
	funcs := c.HandleEndEventFuncs[id]
	funcs = append(funcs, cb)
	c.HandleEndEventFuncs[id] = funcs
}

func (c *CompositeElementHandler) AddGlobalHandleElementFunc(cb HandleElementFunc) {
	c.GlobalHandleElementFuncs = append(c.GlobalHandleElementFuncs, cb)
}

func (c *CompositeElementHandler) AddGlobalHandleStartEventFunc(cb HandleStartEventFunc) {
	c.GlobalHandleStartEventFuncs = append(c.GlobalHandleStartEventFuncs, cb)
}

func (c *CompositeElementHandler) AddGlobalHandleSequenceFlowFunc(cb HandleSequenceFlowFunc) {
	c.GlobalHandleSequenceFlowFuncs = append(c.GlobalHandleSequenceFlowFuncs, cb)
}

func (c *CompositeElementHandler) AddGlobalHandleTaskFunc(cb HandleTaskFunc) {
	c.GlobalHandleTaskFuncs = append(c.GlobalHandleTaskFuncs, cb)
}

func (c *CompositeElementHandler) AddGlobalHandleEndEventFunc(cb HandleEndEventFunc) {
	c.GlobalHandleEndEventFuncs = append(c.GlobalHandleEndEventFuncs, cb)
}

func (c *CompositeElementHandler) AddHook(hook Hook) {
	c.Hooks = append(c.Hooks, hook)
}
