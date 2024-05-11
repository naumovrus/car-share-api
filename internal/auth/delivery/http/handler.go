package http

import "car-api/internal/auth"

type handler struct {
	uc auth.UC
}

func New(uc auth.UC) auth.Handler {
	return &handler{uc: uc}
}
