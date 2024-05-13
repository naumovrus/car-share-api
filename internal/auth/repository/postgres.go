package repository

import (
	"car-api/internal/auth"
	authModel "car-api/internal/auth/model"
	"car-api/internal/entity"
	storage_postgres "car-api/pkg/storage"
)

type repo struct {
	psql storage_postgres.Postgres
}

func New(psql storage_postgres.Postgres) auth.PgRepo {
	return &repo{psql: psql}
}

func (r *repo) CreateUserCreds(params *entity.UserDto) error {
	query := `insert into "user" (email, password, uuid)`
	_, err := r.psql.Exec(query, params.Email, params.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) CreateUserPersonalData(params authModel.UserPersonalInfo, personalDataByte []byte) error {
	query := `insert into user_credentials (client_uuid, is_verifyed, data) values ((select uuid from "user" where email=$1), false, $2)`
	_, err := r.psql.Exec(query, params.Email, params)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) UpdateUserCreds() error {
	panic("IMPLEMENT ME IF NECCESSARY")
}

func (r *repo) GetUserByCreds(params authModel.SignInParams) (*entity.Client, error) {
	var user entity.Client
	query := `select id, email, uuid from "user" where email=$1 and password=$2`
	err := r.psql.Get(&user, query, params.Email, params.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) SetValid(email string) error {
	query := `update user_credentials set is_verifyed=true where client_uuid=(select uuid from "user" where email=$1)`
	_, err := r.psql.Exec(query, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) SaveAuthToken(params authModel.SaveSessionParams) error {
	query := `insert into session (user_id, session_id) values ($1, $2)`
	_, err := r.psql.Exec(query, params.ClientId, params.SessionId)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetAllClient() ([]entity.Client, error) {
	var response []entity.Client
	err := r.psql.Select(&response, `select id, email, u.client_uuid, uc.data from "user" u inner join user_credentials on client_uuid=uc.client_uuid`)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *repo) IsAdmin(uuid string) (bool, error) {
	var roleId int64
	err := r.psql.Get(&roleId, `select role_id from "user" where uuid=$1`, uuid)
	if err != nil {
		return false, err
	}
	return roleId == 1, nil
}
