package engine

import (
	"github.com/ai.zelenin/bpmn/pkg/bpmn/model"
	"sync"
)

type InMemBPMNStorage struct {
	m *sync.Map
}

func NewInMemBPMNStorage() *InMemBPMNStorage {
	return &InMemBPMNStorage{
		m: new(sync.Map),
	}
}

func (i *InMemBPMNStorage) StoreBPMN(id string, bpmn *model.BPMN) error {
	i.m.Store(id, bpmn)
	return nil
}

func (i *InMemBPMNStorage) LoadBPMN(id string) (*model.BPMN, error) {
	bpmn, ok := i.m.Load(id)
	if !ok {
		return nil, ErrBPMNSchemaNotFound
	}
	r, ok := bpmn.(*model.BPMN)
	if !ok {
		return nil, ErrUnexpectedType
	}
	return r, nil
}
