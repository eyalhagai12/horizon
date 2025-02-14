package workflows

import (
	"horizon/server"

	"go.temporal.io/sdk/workflow"
)

type WorkflowFunc[T any, R any] func(workflow.Context, server.Env, T) (R, error)

func FromFunc[T any, R any](wf WorkflowFunc[T, R], env server.Env) func(workflow.Context, T) (R, error) {
	return func(ctx workflow.Context, request T) (R, error) {
		return wf(ctx, env, request)
	}
}
