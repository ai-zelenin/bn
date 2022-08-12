package model

type Node struct {
	Element
	Incoming []*Node
	Outgoing []*Node
}

func NewNode(el Element) *Node {
	return &Node{
		Element:  el,
		Incoming: make([]*Node, 0),
		Outgoing: make([]*Node, 0),
	}
}

func (b *Node) Out() []*Node {
	return b.Outgoing
}

func (b *Node) AddOut(nodes ...*Node) {
	for _, el := range nodes {
		if el != nil {
			b.Outgoing = append(b.Outgoing, el)
		}
	}
}

func (b *Node) In() []*Node {
	return b.Incoming
}

func (b *Node) AddIn(nodes ...*Node) {
	for _, el := range nodes {
		if el != nil {
			b.Incoming = append(b.Outgoing, el)
		}
	}
}

func (b *Node) WalkForward(cb func(node *Node) bool) {
	if !cb(b) {
		return
	}
	outgoing := b.Out()
	for _, out := range outgoing {
		out.WalkForward(cb)
	}
}
