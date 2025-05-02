package dbconn

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Data struct{
	Name string
	Email string
}

func Workflow(ctx workflow.Context,values Data)(string,error){
	ao:=workflow.ActivityOptions{
		StartToCloseTimeout: 10*time.Second,
	}
	ctx=workflow.WithActivityOptions(ctx,ao)
	var result string
	for i:=0;i<20;i++{
	err :=workflow.ExecuteActivity(ctx,Activity,values).Get(ctx,&result)
	
	if err!=nil{
		return "",err
	}
	result+=fmt.Sprintf("iteration %d: %s\n",i+1,result)
}
	return result,nil

}