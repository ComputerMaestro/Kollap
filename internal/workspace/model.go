package workspace

import (
	"time"
)

type Workspace struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Name: name,
	}
}
