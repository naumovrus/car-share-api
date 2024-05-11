package auth

import (
	authModel "car-api/internal/auth/model"
)

type UC interface {
	SignIn(params *authModel.SignInParams) (*authModel.SignInResponse, error)
	SignUp(params *authModel.SignUpParams) error
	AddUserPersonalInfo(params *authModel.UserPersonalInfo) error
	PassportVerify(params *authModel.PassportVerifyParams) error
	GetUserInfo(clientSessionToken string) (*authModel.TotalUserInfoParams, error)
}
