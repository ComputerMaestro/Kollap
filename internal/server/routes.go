package server

import (
	"Collap/internal/workspace"

	"github.com/go-chi/chi/v5"
)

func mainRouting(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		documentRoutes(r)
		workspaceRoutes(r)
	})
}

func documentRoutes(r chi.Router) {
}

func workspaceRoutes(r chi.Router) {
	r.Route("/workspace", func(r chi.Router) {
		r.Post("/create", workspace.CreateWorkspace)
	})
}
