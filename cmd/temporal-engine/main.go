package main

import (
	"context"
	"fmt"
	"github.com/ai.zelenin/bpmn/pkg/engine"
	"go.temporal.io/sdk/activity"
)

func main() {
	err := Start()
	if err != nil {
		panic(err)
	}
}

func Start() error {
	e, err := engine.NewTemporalEngine()
	if err != nil {
		return err
	}
	err = e.Register(&SampleActivities{})
	if err != nil {
		return err
	}

	return nil
}

type SampleActivities struct {
}

func (a *SampleActivities) SampleActivity1(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity2(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity3(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity4(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity5(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + name, nil
}
