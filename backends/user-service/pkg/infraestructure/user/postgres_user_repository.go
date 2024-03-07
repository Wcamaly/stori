package user

import (
	"context"
	"stori/user-service/pkg/config/errors"
	"stori/user-service/pkg/domain/models"
	"stori/user-service/pkg/domain/user"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var _ user.Repository = (*UserPostgressRepository)(nil)

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
	ds := dialect.From(goqu.T("tgf_service.user_schema").As("u")).
		Select(
			goqu.L("distinct c.id"),
			"username",
			"password",
			"lasLogin")

	exprs := exp.NewExpressionList(exp.AndType)

	/* --------------------- User reference filters --------------------- */
	if len(filter.ID) > 0 {
		exprs = exprs.Append(goqu.L(`id`).Eq(filter.ID))
	}

	if len(filter.Username) > 0 {
		exprs = exprs.Append(goqu.L(`username`).Eq(filter.Username))
	}

	sqlQuery, args, err := ds.Prepared(true).Where(exprs).ToSQL()
	if err != nil {
		return nil, err
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
}

func (u UserPostgressRepository) FindById(ctx context.Context, id models.ID) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgressRepository) Create(ctx context.Context, us *user.User) error {
	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.New(user.ErrorUserInternal, "error beginning transaction User")
	}
	_, err = tx.NamedExecContext(ctx, `
		INSERT TO user_schema.user(
		id,username,password
		) VALUES(:id, :username, :password)`,
		map[string]interface{}{
			"id":       us.Id(),
			"username": us.UserName(),
			"pasword":  us.Password(),
		})
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u UserPostgressRepository) Update(ctx context.Context, c *user.User) error {
	//TODO implement me
	panic("implement me")
}
