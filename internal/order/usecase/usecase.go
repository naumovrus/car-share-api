package usecase

import "car-api/internal/order"

type uc struct {
	
}

func New() order.UC {
	return &uc{}
}
