package postgres

import (
	"Collap/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DocumentRepositoryDao domain.DocumentRepository
)

func InitializePostgresPools(dbPool *pgxpool.Pool) {
	DocumentRepositoryDao = NewDocumentRepositoryImpl(dbPool)
}
