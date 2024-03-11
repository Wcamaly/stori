package user

import (
	"context"
	"stori/user-service/pkg/domain/models"
	"stori/user-service/pkg/domain/user"
)

type GetUserRequestById struct {
	ID string `json:"id"`
}

type GetUserByIdDtoResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	SurName   string `json:"surname"`
}

type GetUserById struct {
	repository user.UserRepository
}

func NewGetUserById(
	repository user.UserRepository,
) *GetUserById {
	return &GetUserById{
		repository: repository,
	}
}

func (g *GetUserById) Exec(ctx context.Context, payload *GetUserRequestById) (*GetUserDtoResponse, error) {
	user, err := g.repository.FindById(ctx, models.ID(payload.ID))
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
