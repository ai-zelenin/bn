package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/pb"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"log"
	"math/rand"
)

var mode string

func main() {
	flag.StringVar(&mode, "mode", "w", "")
	flag.Parse()
	switch mode {
	case "w":

	case "l":

	}
	client, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         "localhost:26500",
		UsePlaintextConnection: true,
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	response, err := client.NewTopologyCommand().Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())
	topology, err := client.NewTopologyCommand().Send(context.Background())
	if err != nil {
		panic(err)
	}
	for _, broker := range topology.Brokers {
		fmt.Println("Broker", broker.Host, ":", broker.Port)
		for _, partition := range broker.Partitions {
			fmt.Println("  Partition", partition.PartitionId, ":", roleToString(partition.Role))
		}
	}
	bpmnFile := "./res/bpmn/lvl1.bpmn"
	cmd := client.NewDeployResourceCommand().AddResourceFile(bpmnFile)
	deployResponse, err := cmd.Send(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(deployResponse.String())

	//processID := "Process_1"
	//ExecProcess(client, processID, 1)

	jobWorker := client.NewJobWorker().JobType("T1").Handler(handleJob).Open()

	<-ctx.Done()
	fmt.Println("exit")
	jobWorker.Close()
	jobWorker.AwaitClose()
}

func roleToString(role pb.Partition_PartitionBrokerRole) string {
	switch role {
	case pb.Partition_LEADER:
		return "Leader"
	case pb.Partition_FOLLOWER:
		return "Follower"
	default:
		return "Unknown"
	}
}

func ExecProcess(client zbc.Client, processID string, i int) {
	variables := make(map[string]interface{})
	variables["orderId"] = i
	request, err := client.NewCreateInstanceCommand().BPMNProcessId(processID).LatestVersion().VariablesFromMap(variables)
	if err != nil {
		panic(err)
	}
	result, err := request.Send(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func handleJob(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()
	variables, err := job.GetVariablesAsMap()
	if err != nil {
		failJob(client, job)
		return
	}
	variables["data"] = rand.Intn(10000)
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		failJob(client, job)
		return
	}
	_, err = request.Send(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("Successfully completed job")
}

func failJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())

	ctx := context.Background()
	_, err := client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send(ctx)
	if err != nil {
		panic(err)
	}
}
