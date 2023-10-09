package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgdraganov/super-librarian/internal/model"
)

var errMissingEnvVariable error = errors.New("environment variable not found")

// NewServiceConfig is a constructor function for the ServiceConfig type
func NewServiceConfig() (model.ServiceConfig, error) {
	appEnv, ok := os.LookupEnv("APP_ENV")
	if !ok {
		return model.ServiceConfig{}, fmt.Errorf("%w: APP_ENV", errMissingEnvVariable)
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		return model.ServiceConfig{}, fmt.Errorf("%w: HOST", errMissingEnvVariable)
	}

	httpPort, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		return model.ServiceConfig{}, fmt.Errorf("%w: HTTP_PORT", errMissingEnvVariable)
	}

	// DB config
	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		return model.ServiceConfig{}, fmt.Errorf("%w: DB_USER", errMissingEnvVariable)
	}

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return model.ServiceConfig{}, fmt.Errorf("%w: DB_PASSWORD", errMissingEnvVariable)
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return model.ServiceConfig{}, fmt.Errorf("%w: DB_HOST", errMissingEnvVariable)
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return model.ServiceConfig{}, fmt.Errorf("%w: DB_PORT", errMissingEnvVariable)
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return model.ServiceConfig{}, fmt.Errorf("%w: DB_NAME", errMissingEnvVariable)
	}

	mysqlConfig := model.MySqlConfig{
		Username: dbUser,
		Password: dbPassword,
		Host:     dbHost,
		Port:     dbPort,
		DbName:   dbName,
	}

	return model.ServiceConfig{
		AppEnv:      appEnv,
		HttpPort:    httpPort,
		Host:        host,
		MySqlConfig: mysqlConfig,
	}, nil
}
