package healthcore

import (
	"cryptchat/internal/auth/right"
	"cryptchat/internal/health/core/model"
	"errors"
)

func GetHealth() (healthcoremodel.HealthResponse, error) {
	healthStatuses := []healthcoremodel.HealthStatus{}
	healthStatuses = append(healthStatuses, getDatabaseHealth())

	isErr := false
	for _, status := range healthStatuses {
		if status.Status == "down" {
			isErr = true
		}
	}

	if isErr {
		return healthcoremodel.HealthResponse{HealthStatuses: healthStatuses}, errors.New("")
	}
	return healthcoremodel.HealthResponse{HealthStatuses: healthStatuses}, nil
}

func getDatabaseHealth() healthcoremodel.HealthStatus {
	status := healthcoremodel.HealthStatus{
		System: "database",
		Status: "up",
	}
	err := authright.GetAuthRepository().Ping()
	if err != nil {
		status.Status = "down"
	}
	return status
}
