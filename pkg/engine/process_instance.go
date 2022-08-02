package engine

import (
	"time"
)

type ProcessInstanceState string

const (
	ProcessInstanceREADY ProcessInstanceState = "READY"
)

type ProcessInstance struct {
	ID                  string
	Name                string
	ProcessDefinitionID string
	State               ProcessInstanceState
	CreatedAt           time.Time
	Execution           *Execution
}
