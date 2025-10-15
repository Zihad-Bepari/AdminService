package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBcofig struct {
	Host          string
	Port          string
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	Httpport     string
	Jwtsecretkey string
	DB           *DBcofig
}

func loadconfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)
	}

	servicename := os.Getenv("SERVICE_NAME")
	if servicename == "" {
		fmt.Println("servicename is required")
		os.Exit(1)
	}

	httpport := os.Getenv("HTTP_PORT")
	if httpport == "" {
		fmt.Println("httpport is required")
		os.Exit(1)
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		fmt.Println("host is required")
		os.Exit(1)
	}

	dbport := os.Getenv("DB_PORT")
	if dbport == "" {
		fmt.Println("port is required")
		os.Exit(1)
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		fmt.Println("name is required")
		os.Exit(1)
	}

	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
		fmt.Println("user is required")
		os.Exit(1)
	}

	dbpassword := os.Getenv("DB_PASSWORD")
	if dbpassword == "" {
		fmt.Println("password is required")
		os.Exit(1)
	}

	dbENABLE_SSL_MODE := os.Getenv("DB_ENABLE_SSL_MODE")
	enblSSLMode, err := strconv.ParseBool(dbENABLE_SSL_MODE)

	if err != nil {
		fmt.Println("Invalid enable ssl mode value", err)
		os.Exit(1)
	}

	dbconfig := &DBcofig{
		Host:          host,
		Port:          dbport,
		Name:          dbname,
		User:          dbuser,
		Password:      dbpassword,
		EnableSSLMODE: enblSSLMode,
	}

	configurations = &Config{
		Version:     version,
		ServiceName: servicename,
		Httpport:    httpport,
		DB:          dbconfig,
	}

}

func Getconfig() *Config {
	if configurations == nil {
		loadconfig()
	}
	return configurations
}
