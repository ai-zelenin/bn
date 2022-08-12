package engine

import (
	"context"
	"github.com/ai.zelenin/bpmn/pkg/bpmn/model"
)

type Model interface {
	LoadExecutableModel(ctx context.Context, processID string) (model.Executable, error)
}
