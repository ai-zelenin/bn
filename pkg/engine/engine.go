package engine

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Engine interface {
	CreateProcessInstance(ctx context.Context, processDefinitionID, name string, vars map[string]any) (*ProcessInstance, error)
	Run(ctx context.Context, instance *ProcessInstance, vars map[string]any) error
}

type engine struct {
	BPMNStorage    BPMNStorage
	StorageFactory StorageFactory
}

func (e *engine) Run(ctx context.Context, instance *ProcessInstance) error {
	//exec := instance.Execution
	//storage := exec.Storage
	//processDef := exec.ProcessDefinition

	return nil
}

func (e *engine) CreateProcessInstance(ctx context.Context, processDefinitionID, name string, vars map[string]any) (*ProcessInstance, error) {
	processDef, err := e.BPMNStorage.LoadProcessDefinition(ctx, processDefinitionID)
	if err != nil {
		return nil, err
	}
	instanceID := fmt.Sprintf("%d", rand.Int63())
	execId := fmt.Sprintf("%d", rand.Int63())
	storage, err := e.StorageFactory.CreateTransactionalStorage(ctx, instanceID)
	if err != nil {
		return nil, err
	}
	exec := &Execution{
		ID:                execId,
		ProcessDefinition: processDef,
		Storage:           storage,
	}
	p := &ProcessInstance{
		ID:                  instanceID,
		ProcessDefinitionID: processDefinitionID,
		Name:                name,
		State:               ProcessInstanceREADY,
		CreatedAt:           time.Now(),
		Execution:           exec,
	}
	tx := func(s Storage) error {
		err = storage.StoreProcessInstance(ctx, instanceID, p)
		if err != nil {
			return err
		}
		for key, val := range vars {
			err = storage.Store(ctx, key, val)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err = storage.InTransaction(tx)
	if err != nil {
		return nil, err
	}
	return p, nil
}
