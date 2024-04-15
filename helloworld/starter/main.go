package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"github.com/temporalio/samples-go/helloworld"
        "go.temporal.io/api/enums/v1"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello_world_workflowID",
		TaskQueue: "hello-world",
                WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
                WorkflowExecutionErrorWhenAlreadyStarted: false,
	}

	we1, err := c.ExecuteWorkflow(context.Background(), workflowOptions, helloworld.Workflow, "Temporal1")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we1.GetID(), "RunID", we1.GetRunID())

	we2, err := c.ExecuteWorkflow(context.Background(), workflowOptions, helloworld.Workflow, "Temporal2")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we2.GetID(), "RunID", we2.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we1.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
