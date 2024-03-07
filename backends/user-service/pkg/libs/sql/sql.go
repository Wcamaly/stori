package sql

import (
	"fmt"
	config "stori/user-service/pkg/config/common"

	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func NewDBConfig() DBConfig {
	return DBConfig{
		Username: config.GetEnv("DB_USER", ""),
		Password: config.GetEnv("DB_PASSWORD", ""),
		Host:     config.GetEnv("DB_HOST", ""),
		Name:     config.GetEnv("DB_NAME", ""),
		Port:     config.GetEnv("DB_PORT", ""),
	}
}

func NewDB(config DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	client, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return client, nil
}
