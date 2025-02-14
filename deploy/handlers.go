package deploy

import (
	"horizon/handlers"
	"horizon/server"
	"horizon/workflows"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.temporal.io/sdk/client"
)

func DeployHandler(c echo.Context, env server.Env, req DeployRequest) (DeployResponse, error) {
	ctx := c.Request().Context()
	deployment := NewDeployment(req.Name)

	opt := client.StartWorkflowOptions{
		ID:        "deploy-" + deployment.Name + "-" + deployment.ID.String(),
		TaskQueue: "horizon",
	}
	env.Workflows.ExecuteWorkflow(ctx, opt, workflows.FromFunc(Deploy, env), deployment)

	return DeployResponse{
		Name:   deployment.Name,
		Status: deployment.Status,
	}, nil
}

func RegisterRoutes(api *echo.Group, env server.Env) {
	api.POST("/deploy", handlers.FromFunc(DeployHandler, env, http.StatusAccepted))
}
