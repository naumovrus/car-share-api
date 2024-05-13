package entity

import "time"

type UserDto struct {
	Id           int64       `json:"id" db:"id"`
	Email        string      `json:"email"`
	Role         string      `json:"role"`
	Password     string      `json:"-"`
	Uuid         string      `json:"uuid"`
	PersonalData interface{} `json:"user_personal_data"`
}
type Client struct {
	Id         int64   `json:"id" db:"id"`
	Email      string  `json:"email" db:"email"`
	ClientUUID string  `json:"uuid" db:"uuid"`
	Balance    float64 `json:"balance"`
	Data       []byte  `json:"data"`
}

type Order struct {
	UUID       string    `json:"uuid" db:"uuid"`
	Amount     float64   `json:"amount" db:"amount"`
	ClientUUID string    `json:"client_uuid" db:"client_uuid"`
	CarUUID    string    `json:"car_uuid" db:"car_uuid"`
	DateCreate time.Time `json:"date_create" db:"date_create"`
	StatusId   int64     `json:"status_id" db:"status_id"`
}

type Car struct {
	CarUUID string `json:"car_uuid" db:"car_uuid"`
	Data    []byte `json:"data" db:"data"`
}
