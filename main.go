package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"go-learn/tutorial"

	_ "github.com/lib/pq"
)

func main() {
	if err := createEmployee(); err != nil {
		log.Fatalln(err)
	}
}

func createDBConnection() (context.Context, *sql.DB, error) {
	ctx := context.Background()
	conn, err := sql.Open("postgres", "postgresql://postgres:postgres@postgres:5432/employee?sslmode=disable")
	if err != nil {
		return nil, nil, err
	}
	return ctx, conn, err
}

func createEmployee() error {
	ctx, conn, err := createDBConnection()
	if err != nil {
		return nil
	}
	defer conn.Close()

	queries := tutorial.New(conn)

	employee, err := queries.CreateEmployee(ctx, "𠮷野家")
	if err != nil {
		return err
	}
	fmt.Println(employee)
	return nil
}
