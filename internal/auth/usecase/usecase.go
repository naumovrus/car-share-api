package usecase

import (
	"car-api/internal/auth"
	authModel "car-api/internal/auth/model"
)

type uc struct {
	repo auth.PgRepo
}

func New(repo auth.PgRepo) auth.UC {
	return &uc{
		repo: repo,
	}
}

func (u *uc) SignIn(params *authModel.SignInParams)  {

}

func (u *uc) SignUp(params *authModel.SignUpParams) {

}

func (u *uc) PassportVerify(params *authModel.)
