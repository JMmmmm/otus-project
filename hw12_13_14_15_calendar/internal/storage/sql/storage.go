package sqlstorage

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/stdlib" //nolint
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	DB  *sqlx.DB
	Ctx *context.Context
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, dsn string) (err error) {
	s.Ctx = &ctx

	s.DB, err = sqlx.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("cannot open pgx driver: %w", err)
	}

	return s.DB.PingContext(ctx)
}

func (s *Storage) Close(ctx context.Context) error {
	return s.DB.Close()
}
