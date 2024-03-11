package user

import (
	"context"
	"stori/user-service/pkg/config/errors"
	"stori/user-service/pkg/domain/models"
)

var (
	ErrorUserInternal = errors.Define("user.internal_error")
	ErrorUserExist    = errors.Define("user.user_exist")
)

type UserRepository interface {
	Count(ctx context.Context, filter *UserFilter) (int, error)
	Find(ctx context.Context, filter *UserFilter) ([]*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error) // Fixed the syntax problem here
	FindById(ctx context.Context, id models.ID) (*User, error)
	Create(ctx context.Context, c *User) error
	Update(ctx context.Context, c *User) error
}

type User struct {
	ID        models.ID
	email     string
	firstName string
	surname   string
}

func NewUser(id models.ID, email string, firstName string, surname string) *User {
	return &User{
		ID: id, email: email, firstName: firstName, surname: surname,
	}
}

func (u *User) Id() models.ID {
	return u.ID
}

func (u *User) Email() string {
	return u.email
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) Surname() string {
	return u.surname
}
