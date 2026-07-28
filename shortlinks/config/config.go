package config

import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Templates struct {
		Path string `yaml:"path"`
	} `yaml:"templates"`
	Database struct {
		Prod struct {
			Host     string `yaml:"host"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbname"`
			Port     int    `yaml:"port"`
			SSLMode  string `yaml:"sslmode"`
			TimeZone string `yaml:"timezone"`
		} `yaml:"prod"`
		Test struct {
			Driver string `yaml:"driver"`
			DSN    string `yaml:"dsn"`
		} `yaml:"test"`
	} `yaml:"database"`
}

func VarAsInt(variable string, name string) int {
	v, err := strconv.Atoi(variable)
    if err != nil {
		message := fmt.Sprintf("The %s environment variable type is invalid", name)
		panic(message)
    }
	return v
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	configFile := os.Getenv("CONFIG_FILE")
	log.Printf("Using config file: %v", configFile)
	if configFile != "" && strings.ToUpper(configFile) != "USE_VARS" {
		log.Printf("Using config file: %v", configFile)
	    yamlFile, err := ioutil.ReadFile(configFile)
	    if err != nil {
		    return nil, err
	    }
	    err = yaml.Unmarshal(yamlFile, config)
	    if err != nil {
		    return nil, err
	    }
	} else if strings.ToUpper(configFile) == "USE_VARS" {
		serverHost := os.Getenv("SERVER_HOST")
		serverPortStr := os.Getenv("SERVER_PORT")
		templatesPath := os.Getenv("TEMPLATES_PATH")
		databaseHost := os.Getenv("DATABASE_HOST")
		databasePortStr := os.Getenv("DATABASE_PORT")
		databaseUser := os.Getenv("DATABASE_USER")
		databasePass := os.Getenv("DATABASE_PASSWORD")
		databaseName := os.Getenv("DATABASE_NAME")
		databaseSsl := os.Getenv("DATABASE_SSL_MODE")
		databaseTz := os.Getenv("DATABASE_TIMEZONE")
		databaseTestDriver := os.Getenv("DATABASE_TEST_DRIVER")
		databaseTestDSN := os.Getenv("DATABASE_TEST_DSN")
		if databaseSsl == "" {
			databaseSsl = "disable"
		}
		if databaseTz == "" {
		    databaseTz = "UTC"
		}
		if databaseTestDriver == "" {
			databaseTestDriver = "sqlite"
		}
		if databaseTestDSN == "" {
		    databaseTestDSN = ":memory:"
		}
		serverPort := VarAsInt(serverPortStr, "SERVER_PORT")
		databasePort := VarAsInt(databasePortStr, "DATABASE_PORT")
		config.Server.Host = serverHost
		config.Server.Port = serverPort
		config.Templates.Path = templatesPath
		config.Database.Prod.Host = databaseHost
		config.Database.Prod.Port = databasePort
		config.Database.Prod.User = databaseUser
		config.Database.Prod.Password = databasePass
		config.Database.Prod.DBName = databaseName
		config.Database.Prod.SSLMode = databaseSsl
		config.Database.Prod.TimeZone = databaseTz
		config.Database.Test.Driver = databaseTestDriver
		config.Database.Test.DSN = databaseTestDSN
	} else {
		panic("No config load option was selected, set the USE_VARS or CONFIG_FILE environment variable!")
	}
	return config, nil
}
