package user

import (
	"context"
	"stori/user-service/pkg/domain/user"
)

type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

type GetUserDtoResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	SurName   string `json:"surname"`
}

type GetUserByEmail struct {
	repository user.UserRepository
}

func NewGetUserByEmail(
	repository user.UserRepository,
) *GetUserByEmail {
	return &GetUserByEmail{
		repository: repository,
	}
}

func (g *GetUserByEmail) Exec(ctx context.Context, payload *GetUserByEmailRequest) (*GetUserDtoResponse, error) {
	user, err := g.repository.FindByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	return &GetUserDtoResponse{
		ID:        string(user.Id()),
		Email:     user.Email(),
		FirstName: user.FirstName(),
		SurName:   user.Surname(),
	}, nil

}
