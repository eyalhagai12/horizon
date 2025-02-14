package deploy

type DeployRequest struct {
	Name string `json:"name"`
}

type DeployResponse struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
