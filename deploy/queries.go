package deploy

import "database/sql"

func StoreDeployment(db *sql.DB, deployment Deployment) (Deployment, error) {
	result := db.QueryRow("INSERT INTO deployments (id, name, status, url) VALUES ($1, $2, $3, $4) RETURNING *;", deployment.ID, deployment.Name, deployment.Status, deployment.URL)
	if result.Err() != nil {
		return Deployment{}, result.Err()
	}

	var newDeployment Deployment
	err := result.Scan(&newDeployment.ID, &newDeployment.Name, &newDeployment.CreatedAt, &newDeployment.UpdatedAt, &newDeployment.DeletedAt, &newDeployment.Status, &newDeployment.URL)
	if err != nil {
		return Deployment{}, err
	}

	return newDeployment, nil
}
