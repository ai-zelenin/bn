package engine

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"math/rand"
	"time"
)

type TemporalEngine struct {
	tCli           client.Client
	BPMNStorage    BPMNStorage
	StorageFactory StorageFactory
}

func NewTemporalEngine() (*TemporalEngine, error) {
	opts := client.Options{
		HostPort: client.DefaultHostPort,
	}
	cli, err := client.NewLazyClient(opts)
	if err != nil {
		return nil, err
	}
	_, err = cli.CheckHealth(context.Background(), &client.CheckHealthRequest{})
	if err != nil {
		return nil, err
	}
	return &TemporalEngine{
		tCli: cli,
	}, nil
}

func (t *TemporalEngine) Register(activities any) error {
	w := worker.New(t.tCli, "bpmn", worker.Options{})
	w.RegisterWorkflow(t.Workflow)
	w.RegisterActivity(activities)
	return w.Run(worker.InterruptCh())
}

func (t *TemporalEngine) CreateProcessInstance(ctx context.Context, processDefinitionID, name, startID string, vars VariableMap) (*ProcessInstance, error) {
	instanceID := fmt.Sprintf("%d", rand.Int63())
	startOptions := client.StartWorkflowOptions{
		ID: instanceID,
	}
	p := &ProcessInstance{
		ID:        instanceID,
		ProcessID: processDefinitionID,
		Name:      name,
		StartID:   startID,
		CurrentID: startID,
		State:     ProcessInstanceREADY,
		CreatedAt: time.Now(),
	}
	_, err := t.tCli.ExecuteWorkflow(ctx, startOptions, t.Workflow, p, vars)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (t *TemporalEngine) Workflow(ctx workflow.Context, p *ProcessInstance, vars VariableMap) error {
	bpmn, err := t.BPMNStorage.FindByProcessID(nil, p.ProcessID)
	if err != nil {
		return err
	}
	_, err = bpmn.LoadExecutableModel(nil, p.ProcessID)
	if err != nil {
		return err
	}

	return nil
}
