package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DBConfig struct {
	Port         string `json:"port"`
	Host         string `json:"host"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DriverName   string `json:"driverName"`
	DatabaseName string `json:"databaseName"`
}

func ToString(config *DBConfig) string {
	return config.User + ":" +
		config.Password + "@(" +
		config.Host + ")/" +
		config.DatabaseName

}

func LoadConfigDBFromFile() *DBConfig {
	file, err := os.Open("./dbconfig.json")

	if err != nil {
		log.Fatal("Error opening config file", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := DBConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	return &config
}

func LoadConfigFromENV() *DBConfig {
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	driver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")

	return &DBConfig{
		Port:         port,
		Host:         host,
		User:         user,
		Password:     password,
		DriverName:   driver,
		DatabaseName: dbName,
	}
}
