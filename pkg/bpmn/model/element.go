package model

type Element interface {
	ID() string
	OutRef() []string
	InRef() []string
}

type ElementIndex[T Element] map[string]T

func (e ElementIndex[T]) ToSlice() []Element {
	var slice = make([]Element, 0, len(e))
	for _, t := range e {
		slice = append(slice, t)
	}
	return slice
}

func (e ElementIndex[T]) AddElement(el T) {
	e[el.ID()] = el
}

func (e ElementIndex[T]) Resolve(ids ...string) []Element {
	if ids == nil {
		return nil
	}
	els := make([]Element, 0, len(ids))
	for _, id := range ids {
		el, ok := e[id]
		if ok {
			els = append(els, el)
		}
	}
	return els
}

func (e ElementIndex[T]) MakeNodeRecursive(el Element) *Node {
	node := NewNode(el)
	outEls := e.Resolve(el.OutRef()...)
	for _, outEl := range outEls {
		node.AddOut(e.MakeNodeRecursive(outEl))
	}
	return node
}
