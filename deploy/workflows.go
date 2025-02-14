package deploy

import (
	"horizon/server"

	"go.temporal.io/sdk/workflow"
)

func Deploy(ctx workflow.Context, env server.Env, deployment Deployment) (Deployment, error) {
	logger := workflow.GetLogger(ctx)

	logger.Info("Deploying application", "name", deployment.Name, "deployment", deployment)

	deployment.Status = DeploymentStatusDeployed

	deployment, err := StoreDeployment(env.DB, deployment)
	if err != nil {
		return Deployment{}, err
	}

	logger.Info("Application deployed", "name", deployment.Name)

	return deployment, nil
}
