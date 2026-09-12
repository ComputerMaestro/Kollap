package workspace

import "time"

type GetWorkspaceRequest struct {
	ID string `json:"id"`
}

type GetWorkspaceResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateWorkspaceRequest struct {
	Name string `json:"name"`
}
