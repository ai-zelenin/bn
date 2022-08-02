package model

type Element interface {
	ID() string
	Outgoing() []string
	Incoming() []string
}
