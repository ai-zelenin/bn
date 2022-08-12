package engine

import (
	"time"
)

type ProcessInstanceState string

const (
	ProcessInstanceREADY    ProcessInstanceState = "READY"
	ProcessInstanceACTIVE   ProcessInstanceState = "ACTIVE"
	ProcessInstanceINCIDENT ProcessInstanceState = "INCIDENT"
	ProcessInstanceDONE     ProcessInstanceState = "DONE"
)

type ProcessInstance struct {
	ID        string
	Name      string
	ProcessID string
	StartID   string
	CurrentID string
	State     ProcessInstanceState
	CreatedAt time.Time
}
