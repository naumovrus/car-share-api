package model

import (
	"car-api/internal/entity"
	"time"
)

type SignUpParams struct {
	Email    *string          `json:"email"`
	Password *string          `json:"password"`
	Data     UserPersonalInfo `json:"user_personal_info"`
}

type UserPersonalInfo struct {
	Email       *string   `json:"-"`
	FirstName   *string   `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	LastName    *string   `json:"last_name"`
	PhoneNumber *string   `json:"phone_number"`
	Age         float64   `json:"age"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type SignInParams struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
}

type PassportVerifyParams struct {
	Photo []byte `json:"data"`
}

type TotalUserInfoParams struct {
	Personal UserPersonalInfo `json:"personal_info"`
	Internal entity.Client    `json:"internal"`
}

type SaveSessionParams struct {
	SessionId string `json:"session_id"`
	ClientId  int64  `json:"client_id"`
}

type SignInResponse struct {
	Email string `json:"email"`
}
