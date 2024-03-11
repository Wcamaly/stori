package user

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"stori/user-service/pkg/config/errors"
	"stori/user-service/pkg/domain/models"
	"stori/user-service/pkg/domain/user"
)

var (
	ErrDomainConversion = errors.Define("userdto.toDomain_error")
)

type UserDto struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	SurName   string `json:"surname"`
}

type UserPostgres struct {
	ID        string `db:"id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	SurName   string `db:"surname"`
}

func (p *UserPostgres) toDomain() (*user.User, error) {
	id, err := models.NewID(p.ID)
	if err != nil {
		return nil, err
	}
	return user.NewUser(id, p.Email, p.FirstName, p.SurName), nil
}

func (u UserDto) toDomain() (*user.User, error) {
	id, err := models.NewID(u.ID)
	if err != nil {
		return nil, err
	}

	email, err := models.NewRequiredString(u.Email)
	if err != nil {
		return nil, errors.New(ErrDomainConversion, "missing email")
	}

	return user.NewUser(id, string(email), u.FirstName, u.SurName), nil
}

func fromDomain(u *user.User) UserDto {
	return UserDto{
		ID:        string(u.Id()),
		Email:     u.Email(),
		FirstName: u.FirstName(),
		SurName:   u.Surname(),
	}
}
