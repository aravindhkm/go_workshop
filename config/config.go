package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ApikeyGR string `mapstructure:"APIKEY_GR"`
}

var AppConfig *Config

func LoadConfig() (error) {
	viper.SetConfigFile(".env") // Look for a .env file in the root directory
	viper.AutomaticEnv()        // Read environment variables

	// Attempt to read from .env (optional file)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found or couldn't read it, relying on environment variables")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	AppConfig =  &cfg

	return nil
}
