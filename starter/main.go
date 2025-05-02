package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
	"go-poc/dbconn"
)

func main(){
	c,err:=client.Dial(client.Options{})
	if err!=nil{
		log.Fatalln("unable to create client",err)
	}
	defer c.Close()

	workflowOptions:=client.StartWorkflowOptions{
	
			ID: "go-poc-workflow",
			TaskQueue: "go-poc",	
	}
	values:=dbconn.Data{"kiran","kkr@gmail.com"}
	we,err:=c.ExecuteWorkflow(context.Background(),workflowOptions,dbconn.Workflow,values)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	var result string
	
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}