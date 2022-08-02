package engine

import "github.com/ai.zelenin/bpmn/pkg/model"

type Execution struct {
	ID                string
	ProcessDefinition *model.Process
	Storage           TransactionalStorage
}
