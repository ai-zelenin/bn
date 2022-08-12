package engine

import (
	"context"
	"github.com/ai.zelenin/bpmn/pkg/bpmn/model"
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
	ScopeStorage
}

type ProcessInstanceStorage interface {
	StoreProcessInstance(ctx context.Context, id string, instance *ProcessInstance) error
	LoadProcessInstance(ctx context.Context, id string) (*ProcessInstance, error)
}

type ScopeStorage interface {
	Store(ctx context.Context, variables ...*Variable) error
	Load(ctx context.Context, keys ...string) (vars Variables, err error)
}

type BPMNStorage interface {
	StoreBPMN(ctx context.Context, id string, bpmn *model.BPMN) error
	LoadBPMN(ctx context.Context, id string) (*model.BPMN, error)
	FindByProcessID(ctx context.Context, processID string) (*model.BPMN, error)
}
