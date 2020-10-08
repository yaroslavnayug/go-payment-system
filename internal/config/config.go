package config

import (
	"os"
)

type Config struct {
	PostgresConfig struct {
		HostString string
	}
}

func Read() Config {
	postgresConnectionString := os.Getenv("POSTGRESQL_URL")
	if postgresConnectionString == "" {
		panic("env POSTGRESQL_URL not set")
	}

	config := Config{}
	config.PostgresConfig.HostString = postgresConnectionString
	return config
}
