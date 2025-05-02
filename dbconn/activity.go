package dbconn

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq" 
)

const(
	host="localhost"
	port=5433
	user="postgres"
	password="password"
	dbname="test"
)

func Activity(context context.Context,values Data)(string,error){

		plsqlconn:=fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s ",
	
			host,port,user,password,dbname)
		db,err:=sql.Open("postgres",plsqlconn)
		if err!=nil{
			log.Fatal(err)
		}
		defer db.Close()
		insertStmt:=`INSERT INTO sample(name,email) VALUES($1,$2)`
		_,err=db.Exec(insertStmt,values.Name,values.Email)
		if err!=nil{
			log.Fatal(err)
		}
	return fmt.Sprintf("Inserted data"),nil
}