package engine

import (
	"github.com/ai.zelenin/bpmn/pkg/bpmn/model"
)

type Execution struct {
	ID                string
	ProcessInstanceID string
	model.Executable
	Storage TransactionalStorage
}
