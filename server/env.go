package server

import (
	"database/sql"
	"log"

	"go.temporal.io/sdk/client"
)

type Env struct {
	DB        *sql.DB
	Workflows client.Client
}

func NewEnv(db *sql.DB) Env {
	workflowClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal("failed to create a new workflow client - ", err)
	}
	return Env{DB: db, Workflows: workflowClient}
}
