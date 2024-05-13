package order

import "car-api/internal/entity"

type Repo interface {
	CreateOrder(params entity.Order) error
	GetOrder(uuid *string) ([]entity.Order, error)
}
