package model

import "errors"

var (
	ErrInvalidBPMN    = errors.New("invalid bpmn")
	ErrUnexpectedType = errors.New("unexpected types")
)
