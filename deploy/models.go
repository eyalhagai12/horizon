package deploy

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const (
	DeploymentStatusPending  = "pending"
	DeploymentStatusRunning  = "running"
	DeploymentStatusFailed   = "failed"
	DeploymentStatusDeployed = "deployed"
)

type Deployment struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Status    string
	URL       string
}

func NewDeployment(name string) Deployment {
	deploymentUuid := uuid.New()
	return Deployment{
		ID:     deploymentUuid,
		Name:   name,
		Status: "pending",
		URL:    "http://localhost:8080/deployments/" + deploymentUuid.String(),
	}
}
