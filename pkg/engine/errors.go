package engine

import "errors"

var (
	ErrUnexpectedType     = errors.New("unexpected type")
	ErrBPMNSchemaNotFound = errors.New("BPMN schema not found")
)
