package repository

import (
	"car-api/internal/entity"
	"car-api/internal/order"
	storage_postgres "car-api/pkg/storage"
)

type repo struct {
	psql storage_postgres.Postgres
}

func New(psql storage_postgres.Postgres) order.Repo {
	return &repo{psql: psql}
}

func (r *repo) CreateOrder(params entity.Order) error {
	query := `insert into order (uuid, car_uuid, client_uuid, status_id, amount) values ($1, $2, $3, $4, $5)`
	_, err := r.psql.Exec(query, params.UUID, params.CarUUID, params.ClientUUID, 1, params.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetOrder(uuid *string) ([]entity.Order, error) {
	var response []entity.Order
	query := `select uuid, car_uuid, client_uuid, status_id, amount, date_create from order where (uuid=$1 or $1 is null) order by date_create desc`
	err := r.psql.Select(&response, query, uuid)
	if err != nil {
		return nil, err
	}
	return response, nil
}
