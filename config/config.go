package config

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// func GetPortEnv() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	port := os.Getenv("APP_PORT")
// 	return port
// }
// func GetDbNameEnv() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	name := os.Getenv("DB_NAME")
// 	return name
// }

// func GetDbPasswordEnv() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	pass := os.Getenv("DB_PASSWORD")
// 	return pass
// }

// func GetDbHostEnv() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	host := os.Getenv("DB_HOST")
// 	return host
// }

// func GetDbUserEnv() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	user := os.Getenv("DB_USER")
// 	return user
// }

// func GetDbPortEnv() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	dbPort := os.Getenv("DB_PORT")
// 	return dbPort
// }

// package configs

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