package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	LogConfig struct {
		Level logrus.Level
	}
	PostgresConfig struct {
		HostString string
	}
}

func GetConfig() Config {
	logLevel := os.Getenv("LOG_LEVEL")
	postgresConnectionString := os.Getenv("POSTGRESQL_URL")

	if logLevel == "" {
		panic("env LOG_LEVEL not set")
	}
	if postgresConnectionString == "" {
		panic("env POSTGRESQL_URL not set")
	}

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		panic(fmt.Sprintf("fail to build logger config: %+v", err))
	}

	config := Config{}
	config.LogConfig.Level = level
	config.PostgresConfig.HostString = postgresConnectionString

	return config
}
