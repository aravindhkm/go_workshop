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
	IsIndex   bool   `json:isIndex`
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

func SetIndexFilter(status bool) {
	AppConfig.IsIndex = true
}

func findFuncKey(dirName string, funcName string) (int, error) {
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
	n, err := strconv.Atoi(funcName)
	if err != nil {
		fmt.Println("Invalid function name:", funcName)
		return 0, err
	}
	if AppConfig.IsReverse {
		result = fileCount - (n + 1)
	} else {
		result = n - 1
	}
	return result, nil
}

func findFileIndex(dirName string, funcName string) (int, error) {
	entries, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	for index, e := range entries {
		fmt.Println("e", index, e.Name(), e.Name() == funcName)
		if e.Name() == funcName {
			return index, nil
		}
	}
	return 0, nil
}

func FindRunIndex(dirName string, funcName string) (int, error) {
	if AppConfig.IsIndex {
		return findFuncKey(dirName, funcName)
	} else {
		return findFileIndex(dirName, funcName)
	}
}
