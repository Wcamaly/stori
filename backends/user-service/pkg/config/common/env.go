package config

import "stori/user-service/pkg/config/errors"

var (
	ErrInvalidEnv = errors.Define("models.invalid_env")
)

type Env string

const (
	Local Env = "local"
	Dev   Env = "dev"
	Stage Env = "stage"
	Prod  Env = "prod"
)

var allowedEnvs = map[string]Env{
	Local.String(): Local,
	Dev.String():   Dev,
	Stage.String(): Stage,
	Prod.String():  Prod,
}

func NewEnv(env string) (Env, error) {
	if env, ok := allowedEnvs[env]; ok {
		return env, nil
	}

	return "", errors.New(
		ErrInvalidEnv,
		"invalid env",
		errors.WithMetadata("env", env),
	)
}

func (e Env) String() string {
	return string(e)
}
