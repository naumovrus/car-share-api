package auth

import (
	authModel "car-api/internal/auth/model"
	"car-api/internal/entity"
)

type PgRepo interface {
	CreateUserCreds(params *entity.UserDto) error
	CreateUserPersonalData(params authModel.UserPersonalInfo, personalDataByte []byte) error
	UpdateUserCreds() error
	GetUserByCreds(params authModel.SignInParams) (*entity.Client, error)
	SetValid(email string) error
	SaveAuthToken(params authModel.SaveSessionParams) error
}

type RedisRepo interface {
}
