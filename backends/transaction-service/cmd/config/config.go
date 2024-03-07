package config

import (
	config "stori/transaction/pkg/config/common"
)

type Config struct {
	Port          string
	CommonConfigs config.Config
}

func LoadConfig() (*Config, error) {

	configs := &Config{
		Port: config.GetEnv("PORT", "8080"),
	}

	if commonConfigs := config.Common(); commonConfigs != nil {
		configs.CommonConfigs = *commonConfigs
	}

	return configs, nil
}
