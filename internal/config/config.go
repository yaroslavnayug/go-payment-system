package config

import (
	"os"

	"github.com/jackc/pgx/v4"
)

type Config struct {
	PostgresConfig struct {
		HostString     string
		MaxConnections int32
		MinConnections int32
		LogLevel       pgx.LogLevel
	}
}

func Read() Config {
	postgresConnectionString := os.Getenv("POSTGRESQL_URL")
	if postgresConnectionString == "" {
		panic("env POSTGRESQL_URL not set")
	}

	config := Config{}
	config.PostgresConfig.HostString = postgresConnectionString

	// TODO: Пробрасывать снаружи
	config.PostgresConfig.MaxConnections = 10
	config.PostgresConfig.MinConnections = 1
	config.PostgresConfig.LogLevel = pgx.LogLevelError

	return config
}
