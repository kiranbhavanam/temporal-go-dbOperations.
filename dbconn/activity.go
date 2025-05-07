package dbconn

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "password"
	dbname   = "test"
)

func Activity(ctx context.Context, values Data) (string, error) {
	// Fix connection string: use dbname instead of database and add sslmode
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return "", fmt.Errorf("database connection error: %v", err)
	}
	defer db.Close()

	// Test connection
	err = db.Ping()
	if err != nil {
		return "", fmt.Errorf("database ping error: %v", err)
	}

	// insertStmt := `INSERT INTO sample(name, email) VALUES($1, $2)`
	// _, err = db.Exec(insertStmt, values.Name, values.Email)
	// if err != nil {
	// 	return "", fmt.Errorf("insert error: %v", err)
	// }

	
	insert2:=`INSERT INTO newsample DEFAULT VALUES`
	_,err=db.Exec(insert2)
	if err!=nil{
		return "",fmt.Errorf("insert error|: %v",err)
	}
	return `successfully inserted to newsample`,nil
	// return fmt.Sprintf("Successfully inserted data for %s", values.Name), nil
}
