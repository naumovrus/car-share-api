package repository

import (
	"car-api/internal/car"
	"car-api/internal/entity"
	storage_postgres "car-api/pkg/storage"
)

type repo struct {
	psql storage_postgres.Postgres
}

func New(psql storage_postgres.Postgres) car.Repo {
	return &repo{psql: psql}
}

func (r *repo) GetCar(uuid *string) ([]entity.Car, error) {
	var response []entity.Car
	query := `select uuid, data from car where (uuid=$1 or $1 is null)`
	err := r.psql.Select(&response, query, uuid)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *repo) AddCar(params entity.Car) error {
	query := `insert into car (uuid, data) values ($1, $2)`
	_, err := r.psql.Exec(query, params.CarUUID, params.Data)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteCar(uuid *string) error {
	query := `delete from car where uuid=$1`
	_, err := r.psql.Exec(query, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) UpdateCar() {}
