package repository

import (
	"car-api/internal/auth"
	"car-api/internal/entity"
	storage_postgres "car-api/pkg/storage"
)

type repo struct {
	psql storage_postgres.Postgres
}

func New(psql storage_postgres.Postgres) auth.PgRepo {
	return &repo{psql: psql}
}

func (r *repo) CreateUserCreds(params entity.Client) (*string, error) {
	_ = `insert into "user" (email, password)`

	return nil, nil
}

func (r *repo) UpdateUserCreds()
