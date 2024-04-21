package entity

import "time"

type Client struct {
	Id         int64   `json:"id" db:"id"`
	Email      string  `json:"email" db:"email"`
	ClientUUID string  `json:"uuid" db:"uuid"`
	Balance    float64 `json:"balance"`
}

type Order struct {
	UUID       string    `json:"uuid" db:"uuid"`
	Amount     float64   `json:"amount" db:"amount"`
	ClientUUID string    `json:"client_uuid" db:"client_uuid"`
	CarUUID    string    `json:"car_uuid" db:"car_uuid"`
	DateCreate time.Time `json:"date_create" db:"date_create"`
	DateUpdate time.Time `json:"date_update" db:"date_update"`
}

type Car struct {
	CarUUID string `json:"car_uuid" db:"car_uuid"`
	Data    []byte `json:"data" db:"data"`
}
