package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	ApikeyGR  string `json:"APIKEY_GR"`
	IsReverse bool   `json:"isReverse"`
}

var AppConfig *Config

func LoadConfig() error {
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
	AppConfig = &cfg

	return nil
}

func SetReverse(status bool) {
	AppConfig.IsReverse = status
}

func FindFuncKey(dirName string, funcName string) (int, error) {
	fileCount := 0
	entries, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			fileCount++
		}
	}

	var result int = 0

	if AppConfig.IsReverse {
		n, err := strconv.Atoi(funcName)
		if err != nil {
			fmt.Println("Invalid function name:", funcName)
			return 0, err
		}
		result = fileCount - n
	}

	return result, nil
}
