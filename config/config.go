package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	API        API
	DBPostgres Database
}

func InitConfig() (*Config, error) {
	// this init from viper documentation
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	api := API{
		Host: viper.GetString("api.host"),
		Port: viper.GetInt("api.port"),
	}

	// check every var of api
	api.validateAPI()

	databasePostgres := Database{
		Conn: viper.GetString("database.postgres.conn"),
	}

	// check every var of database postgres
	err = databasePostgres.validateDatabasePostgres()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return &Config{
		API:        api,
		DBPostgres: databasePostgres,
	}, nil
}

func (a *API) validateAPI() {
	if a.Host == "" {
		a.Host = "localhost"
	}

	if a.Port == 0 {
		a.Port = 8000
	}
}

func (d *Database) validateDatabasePostgres() error {
	if d.Conn == "" {
		return errors.New("database connection is nil")
	}

	return nil
}
