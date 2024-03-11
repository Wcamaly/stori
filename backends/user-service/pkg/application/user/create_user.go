package user

import (
	"context"
	"stori/user-service/pkg/config/errors"
	"stori/user-service/pkg/domain/models"
	"stori/user-service/pkg/domain/user"
)

var (
	ErrorHashPassword = errors.Define("create_user.hash_password_error")
)

type CreateUser struct {
	repository user.UserRepository
}

func NewCreateUser(
	repository user.UserRepository,
) *CreateUser {
	return &CreateUser{
		repository: repository,
	}
}

type CreateUserDto struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surName"`
}

type CreateUserResponse struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surName"`
}

func (cu *CreateUser) Exec(ctx context.Context, payload *CreateUserDto) (*CreateUserResponse, error) {

	id := models.GenerateUUID()

	exist, err := cu.repository.FindByEmail(ctx, payload.Email)

	if err != nil {
		return nil, err
	}

	if exist != nil {
		return nil, errors.New(user.ErrorUserExist, "user already exist")
	}

	newUser := user.NewUser(id, payload.Email, payload.FirstName, payload.Surname)
	err = cu.repository.Create(ctx, newUser)
	if err != nil {
		return nil, errors.New(user.ErrorUserInternal, "error creating user")
	}
	return &CreateUserResponse{Id: id.String(), FirstName: payload.FirstName, Surname: payload.Surname, Email: payload.Email}, nil
}
