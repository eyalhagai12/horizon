package workflows

import (
	"horizon/server"

	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/workflow"
)

type EnvInterceptor struct {
	interceptor.WorkerInterceptorBase
	Env server.Env
}

func NewEnvInterceptor(env server.Env) EnvInterceptor {
	return EnvInterceptor{Env: env}
}

func (e EnvInterceptor) Interceptworkflow(ctx workflow.Context, next interceptor.WorkflowInboundInterceptor) interceptor.WorkflowInboundInterceptor {
	return &interceptor.WorkflowInboundInterceptorBase{
		Next: next,
	}
}
