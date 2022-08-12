package engine

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Engine interface {
	CreateProcessInstance(ctx context.Context, processID, name string, vars VariableMap) (*ProcessInstance, error)
}

type engine struct {
	BPMNStorage    BPMNStorage
	StorageFactory StorageFactory
}

func (e *engine) CreateProcessInstance(ctx context.Context, processID, name string, vars VariableMap) (*ProcessInstance, error) {
	instanceID := fmt.Sprintf("%d", rand.Int63())
	storage, err := e.StorageFactory.CreateTransactionalStorage(ctx, instanceID)
	if err != nil {
		return nil, err
	}
	p := &ProcessInstance{
		ID:        instanceID,
		ProcessID: processID,
		Name:      name,
		State:     ProcessInstanceREADY,
		CreatedAt: time.Now(),
	}
	tx := func(s Storage) error {
		err = s.StoreProcessInstance(ctx, instanceID, p)
		if err != nil {
			return err
		}
		err = s.Store(ctx, vars.ToSlice()...)
		if err != nil {
			return err
		}
		return nil
	}
	err = storage.InTransaction(tx)
	if err != nil {
		return nil, err
	}
	return p, nil
}
