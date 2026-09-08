package postgres

import (
	"Collap/internal/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DocumentRepositoryImpl struct {
	DB *pgxpool.Pool
}

func NewDocumentRepositoryImpl(db *pgxpool.Pool) domain.DocumentRepository {
	return &DocumentRepositoryImpl{DB: db}
}

func (r *DocumentRepositoryImpl) Create(ctx context.Context, document *domain.Document) error {
	query := `INSERT INTO documents (name, email) VALUES ($1, $2) RETURNING id`
	return
}
