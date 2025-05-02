package main

import (
	"go-poc/dbconn"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main(){
	c,err:=client.Dial(client.Options{})
	if err!=nil{
		log.Fatal(err)
	}
	defer c.Close()
	w:=worker.New(c,"go-poc",worker.Options{})
	w.RegisterActivity(dbconn.Activity)
	w.RegisterWorkflow(dbconn.Workflow)
	err=w.Run(worker.InterruptCh())
	if err!=nil{
		log.Fatalln("Unable to start the worker",err)
	}
}