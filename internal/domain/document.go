package domain

import "context"

type Document struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type DocumentRepository interface {
	Create(ctx context.Context, user *Document) error
}
