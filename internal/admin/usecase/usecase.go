package admin

import (
	"car-api/internal/admin"
	"car-api/internal/auth"
	"car-api/internal/car"
	"car-api/internal/entity"
	"car-api/internal/order"
	"car-api/pkg/logger"
)

type uc struct {
	orderRepo  order.Repo
	carRepo    car.Repo
	clientRepo auth.PgRepo
	logger     logger.ApiLogger
}

func New(orderRepo order.Repo,
	carRepo car.Repo,
	clientRepo auth.PgRepo,
	logger logger.ApiLogger) admin.UC {
	return &uc{
		orderRepo:  orderRepo,
		carRepo:    carRepo,
		clientRepo: clientRepo,
		logger:     logger,
	}
}

func (u *uc) GetClient(uuid *string) ([]entity.Client, error) {
	
	return nil, nil
}

func (u *uc) GetOrder(uuid *string) ([]entity.Order, error) {
	return nil, nil
}

func (u *uc) GetCar(uuid *string) ([]entity.Car, error) {
	return nil, nil
}
