package main

import (
	"time"

	"github.com/google/uuid"
)

type WorkspaceDto struct {
	Name string `json:"name"`
}

type WorkspaceResponse struct {
	Workspaces []struct {
		ID        uuid.UUID `json:"workspaceId"`
		Name      string    `json:"name"`
		UserCount int       `json:"userCount"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"workspaces"`
}
