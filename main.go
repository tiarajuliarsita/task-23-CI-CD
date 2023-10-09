package main

import (
	"clean_arch/config"
	"clean_arch/router"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	db := config.InitDBMysql(cfg)
	config.InitMigrate(db)
	app := echo.New()
	router.New(app, db)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", cfg.APP_PORT)))
}
