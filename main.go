package main

import (
	"clean_arch/config"
	"clean_arch/router"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	// config.InitDBMysql 	 ()
	cfg := config.InitConfig()
	db := config.InitDBMysql(cfg)
	config.InitMigrate(db)
	app := echo.New()
	router.New(app, db)

	// port := cfg.APP_PORT
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", cfg.APP_PORT)))
}
