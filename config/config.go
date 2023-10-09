package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	APP_PORT    int
	DB_PORT     int
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func InitConfig() *AppConfig {
	var res = new(AppConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *AppConfig {
	var res = new(AppConfig)

	 godotenv.Load(".env")
	
	// var isRead = false
	if val, found := os.LookupEnv("APP_PORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid port value,", err.Error())
			return nil
		}
		// isRead = true
		res.APP_PORT = port
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid db port value,", err.Error())
			return nil
		}
		// isRead = true
		res.DB_PORT = port
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		// isRead = true
		res.DB_HOST = val
	}

	if val, found := os.LookupEnv("DB_USER"); found {
		// isRead = true
		res.DB_USER = val
	}

	if val, found := os.LookupEnv("DB_PASSWORD"); found {
		// isRead = true
		res.DB_PASSWORD = val
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		// isRead = true
		res.DB_NAME = val
	}

	return res

}