package workspace

import (
	"context"

	"gorm.io/gorm"
)

var (
	WorkspaceDao WorkspaceRepository
)

type WorkspaceRepository interface {
	FindByID(ctx context.Context, id string) (*Workspace, error)
	Create(ctx context.Context, w *Workspace) (*Workspace, error)
}

func InitializeWorkspaceRepository(db *gorm.DB) {
	WorkspaceDao = NewWorkspaceRepository(db)
}

type WorkspaceRepositoryImpl struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) *WorkspaceRepositoryImpl {
	return &WorkspaceRepositoryImpl{
		db: db,
	}
}

func (r *WorkspaceRepositoryImpl) FindByID(ctx context.Context, id string) (*Workspace, error) {
	var workspace Workspace
	if err := r.db.WithContext(ctx).First(&workspace, id).Error; err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (r *WorkspaceRepositoryImpl) Create(ctx context.Context, w *Workspace) (*Workspace, error) {
	res := r.db.WithContext(ctx).Create(w)
	if res.Error != nil {
		return nil, res.Error
	}
	return w, nil
}
