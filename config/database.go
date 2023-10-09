package config

import (
	"clean_arch/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	DB  *gorm.DB
)

// func ConnectDB() {

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	port := GetDbPortEnv()
// 	password := GetDbPasswordEnv()
// 	host := GetDbHostEnv()
// 	name := GetDbNameEnv()
// 	user := GetDbUserEnv()

// 	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
// 	DB, err = gorm.Open(mysql.Open(config), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("error connecting to database: ", err)
// 	}
// 	log.Println("connected to database")
// 	initMigrate()

// }

func InitDBMysql(cfg *AppConfig) *gorm.DB {

	// declare struct config & variable connectionString
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	
	return db
}

func InitMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&model.User{})
	
}
