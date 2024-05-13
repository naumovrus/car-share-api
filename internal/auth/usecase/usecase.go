package usecase

import (
	"car-api/internal/auth"
	authModel "car-api/internal/auth/model"
	"car-api/internal/entity"
	"car-api/internal/failure"
	"car-api/pkg/logger"
	"car-api/pkg/validator"
	"fmt"

	gojson "github.com/goccy/go-json"

	"github.com/google/uuid"
)

type uc struct {
	repo   auth.PgRepo
	logger logger.ApiLogger
}

func New(repo auth.PgRepo) auth.UC {
	return &uc{
		repo: repo,
	}
}

func (u *uc) saveAuthToken(params authModel.SaveSessionParams) error {
	err := u.repo.SaveAuthToken(params)
	if err != nil {
		return fmt.Errorf("error while save auth token: %v", err)
	}
	return nil
}

func (u *uc) SignIn(params *authModel.SignInParams) (*authModel.SignInResponse, error) {
	if params == nil {
		return nil, failure.ErrInput
	}
	userObj, err := u.repo.GetUserByCreds(*params)
	if err != nil {
		u.logger.Errorf("error while fetch user: %v", err)
		return nil, fmt.Errorf("user not found")
	}

	return &authModel.SignInResponse{
		Email: userObj.Email,
	}, nil
}

func (u *uc) SignUp(params *authModel.SignUpParams) error {
	clientUUID := uuid.New().String()
	if params.Email == nil || params.Password == nil {
		u.logger.Errorf("error while sign up user: password or email is null")
		return fmt.Errorf("email or password is null")
	}
	err := u.repo.CreateUserCreds(&entity.UserDto{
		Email:    *params.Email,
		Password: *params.Password,
		Uuid:     clientUUID,
	})
	if err != nil {
		u.logger.Errorf("error while create user: %v", err)
		return fmt.Errorf("can't create user")
	}
	return nil
}

func (u *uc) AddUserPersonalInfo(params *authModel.UserPersonalInfo) error {
	if params == nil {
		return fmt.Errorf("data is null")
	}
	if validator.ValidateAge(params.DateOfBirth) {
		return fmt.Errorf("user is too young for a rent")
	}

	personalDataByte, _ := gojson.Marshal(*params)

	err := u.repo.CreateUserPersonalData(*params, personalDataByte)
	if err != nil {
		return fmt.Errorf("can't validate user personal data")
	}
	return nil
}

func (u *uc) PassportVerify(params *authModel.PassportVerifyParams) error {

	return nil
}
func (u *uc) GetUserInfo(clientSessionToken string) (*authModel.TotalUserInfoParams, error) {

	return nil, nil
}
