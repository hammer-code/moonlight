package config

import (
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
		Host:     viper.GetString("api.database.host"),
		Port:     viper.GetInt("api.database.port"),
		Username: viper.GetString("api.database.username"),
		Password: viper.GetString("api.database.password"),
	}

	// check every var of database postgres
	databasePostgres.validateDatabasePostgres()

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

func (a *Database) validateDatabasePostgres() {
	if a.Host == "" {
		a.Host = "localhost"
	}

	if a.Port == 0 {
		a.Port = 8000
	}

	if a.Username == "" {
		a.Username = "postgres"
	}

	if a.Password == "" {
		a.Password = ""
	}

}
