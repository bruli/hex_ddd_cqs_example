package config

import (
	"errors"
	"os"
)

type Config struct {
	postgresDB, postgresUser, postgresPassword string
	postgresHost, postgresPort                 string
	apiHost                                    string
}

func (c Config) PostgresHost() string {
	return c.postgresHost
}

func (c Config) PostgresPort() string {
	return c.postgresPort
}

func (c Config) PostgresDB() string {
	return c.postgresDB
}

func (c Config) PostgresUser() string {
	return c.postgresUser
}

func (c Config) PostgresPassword() string {
	return c.postgresPassword
}

func (c Config) ApiHost() string {
	return c.apiHost
}

func New() (*Config, error) {
	dbName := os.Getenv("POSTGRES_DATABASE")
	if dbName == "" {
		return nil, errors.New("POSTGRES_DATABASE environment variable not set")
	}
	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		return nil, errors.New("POSTGRES_USER environment variable not set")
	}
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbPassword == "" {
		return nil, errors.New("POSTGRES_PASSWORD environment variable not set")
	}
	apiHost := os.Getenv("API_HOST")
	if apiHost == "" {
		return nil, errors.New("API_HOST environment variable not set")
	}
	dbHost := os.Getenv("POSTGRES_HOST")
	if dbHost == "" {
		return nil, errors.New("POSTGRES_HOST environment variable not set")
	}
	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		return nil, errors.New("POSTGRES_PORT environment variable not set")
	}
	return &Config{
		postgresDB:       dbName,
		postgresUser:     dbUser,
		postgresPassword: dbPassword,
		postgresHost:     dbHost,
		postgresPort:     dbPort,
		apiHost:          apiHost,
	}, nil
}
