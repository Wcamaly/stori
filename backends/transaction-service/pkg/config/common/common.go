package config

import (
	"fmt"
	"sync"
)

var once sync.Once
var cfg *Config

type Config struct {
	Env     Env
	Service string
}

func Common() *Config {
	once.Do(func() {
		env, err := NewEnv(GetEnv("ENV", "local"))
		if err != nil {
			panic(fmt.Errorf("invalid env: %s", err))
		}

		cfg = &Config{
			Env:     env,
			Service: GetEnv("SERVICE", "service"),
		}
	})

	return cfg
}
