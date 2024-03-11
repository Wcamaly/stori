package user

import (
	"context"
	"database/sql"
	"stori/user-service/pkg/config/errors"
	"stori/user-service/pkg/domain/models"
	"stori/user-service/pkg/domain/user"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var _ user.UserRepository = (*UserPostgressRepository)(nil)

func NewPostgresUserRepository(db *sqlx.DB) *UserPostgressRepository {
	return &UserPostgressRepository{
		db: db,
	}
}

type UserPostgressRepository struct {
	db *sqlx.DB
}

func (u UserPostgressRepository) Count(ctx context.Context, filter *user.UserFilter) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgressRepository) Find(ctx context.Context, filter *user.UserFilter) ([]*user.User, error) {

	dialect := goqu.Dialect("postgres")
	ds := dialect.From(goqu.T("user").As("u")).
		Select(
			goqu.L("distinct u.id"),
			"u.email",
			"u.first_name",
			"u.surname")

	exprs := exp.NewExpressionList(exp.AndType)

	/* --------------------- User reference filters --------------------- */
	if len(filter.ID) > 0 {
		exprs = exprs.Append(goqu.L(`id`).Eq(filter.ID))
	}

	if len(filter.Email) > 0 {
		exprs = exprs.Append(goqu.L(`email`).Eq(filter.Email))
	}

	sqlQuery, args, err := ds.Prepared(true).Where(exprs).ToSQL()
	if err != nil {
		return nil, err
	}
	type UserFilter struct {
		ID    string
		Email string
	}

	var userPostgres []*UserPostgres

	if err := u.db.SelectContext(
		ctx,
		&userPostgres,
		sqlQuery,
		args...,
	); err != nil {
		return nil, errors.Wrap(
			user.ErrorUserInternal,
			err,
			"contract internal error",
			errors.WithMetadata("filters", filter),
			errors.WithMetadata("query", sqlQuery),
			errors.WithMetadata("args", args),
		)
	}

	users := make([]*user.User, len(userPostgres))
	for i, c := range userPostgres {
		users[i], err = c.toDomain()
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func (u UserPostgressRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {

	var userPostgres UserPostgres
	if err := u.db.GetContext(
		ctx,
		&userPostgres,
		`SELECT distinct id, u.email, u.first_name, u.surname FROM public.user AS u WHERE (u.email = $1)`,
		email,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(
			user.ErrorUserInternal,
			err,
			"contract internal error",
			errors.WithMetadata("email", email),
		)
	}

	return userPostgres.toDomain()

}

func (u UserPostgressRepository) FindById(ctx context.Context, id models.ID) (*user.User, error) {
	/* dialect := goqu.Dialect("postgres")
	ds := dialect.From(goqu.T("user").As("u")).
		Select(
			goqu.L("distinct u.id"),
			"u.email",
			"u.first_name",
			"u.surname").
		Where(goqu.L("u.id").Eq(id))

	sqlQuery, args, err := ds.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	} */

	var userPostgres UserPostgres
	if err := u.db.GetContext(
		ctx,
		&userPostgres,
		`SELECT distinct id, u.email, u.first_name, u.surname FROM public.user AS u WHERE (u.id = $1)`,
		id,
	); err != nil {
		return nil, errors.Wrap(
			user.ErrorUserInternal,
			err,
			"contract internal error",
			errors.WithMetadata("id", id),
			//errors.WithMetadata("query", sqlQuery),
			//errors.WithMetadata("args", args),
		)
	}

	return userPostgres.toDomain()
}

func (u UserPostgressRepository) Create(ctx context.Context, us *user.User) error {
	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.New(user.ErrorUserInternal, "error beginning transaction User")
	}
	_, err = tx.NamedExecContext(ctx, `
		INSERT INTO public.user(
		id, email, first_name, surname
		) VALUES(:id, :email, :first_name, :surname)`,
		map[string]interface{}{
			"id":         us.Id(),
			"email":      us.Email(),
			"first_name": us.FirstName(),
			"surname":    us.Surname(),
		})
	if err != nil {
		println("err: ", err.Error())
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u UserPostgressRepository) Update(ctx context.Context, c *user.User) error {
	//TODO implement me
	panic("implement me")
}
