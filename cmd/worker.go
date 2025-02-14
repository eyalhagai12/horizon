package main

import (
	"database/sql"
	"horizon/deploy"
	"horizon/server"
	"horizon/workflows"
	"log"

	_ "github.com/lib/pq"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func RegisterWorkflows(env server.Env, w worker.Worker) {
	w.RegisterWorkflowWithOptions(workflows.FromFunc(deploy.Deploy, env), workflow.RegisterOptions{Name: "Deploy"})
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=horizon-dev password=horizon-dev sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to the database - ", err)
	}
	env := server.NewEnv(db)

	w := worker.New(env.Workflows, "horizon", worker.Options{})
	RegisterWorkflows(env, w)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal("unable to start worker", err)
	}
}
