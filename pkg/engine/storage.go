package engine

import (
	"context"
	"github.com/ai.zelenin/bpmn/pkg/model"
)

type StorageFactory interface {
	CreateTransactionalStorage(ctx context.Context, instanceID string) (TransactionalStorage, error)
}

type TransactionalFunc func(s Storage) error

func InTransaction(txStore TransactionalStorage, txFunc TransactionalFunc) error {
	err := txStore.Begin()
	if err != nil {
		return err
	}
	err = txFunc(txStore)
	if err != nil {
		return err
	}
	return txStore.Commit()
}

type TransactionalStorage interface {
	Begin() error
	Storage
	Commit() error
	InTransaction(txFunc TransactionalFunc) error
}

type Storage interface {
	ProcessInstanceStorage
	Scope
}

type ProcessInstanceStorage interface {
	StoreProcessInstance(ctx context.Context, id string, instance *ProcessInstance) error
	LoadProcessInstance(ctx context.Context, id string) (*ProcessInstance, error)
}

type Scope interface {
	Store(ctx context.Context, key string, val any) error
	Load(ctx context.Context, key string) (val any, err error)
}

type BPMNStorage interface {
	StoreBPMN(ctx context.Context, id string, bpmn *model.BPMN) error
	LoadBPMN(ctx context.Context, id string) (*model.BPMN, error)
	LoadProcessDefinition(ctx context.Context, id string) (*model.Process, error)
}
