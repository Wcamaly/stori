package user

import (
	"context"
	"stori/user-service/pkg/domain/user"
)

type GetUserByEmail struct {
	email string `json:"email"`
}

type GetUser struct {
	repository user.Repository
}

func NewGetUser(
	repository user.Repository,
) *GetUser {
	return &GetUser{
		repository: repository,
	}
}

type GetUserDtoResponse struct {
	username  string `json:"username"`
	lastLogin string `json:"lastLogin"`
}

func (g *GetUser) Exec(ctx context.Context, payload *GetUserByEmail) (*GetUserDtoResponse, error) {

}
