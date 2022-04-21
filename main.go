package main

import (
	"fmt"
	"user/config"
	user "user/controller"
	mUser "user/model"
	"user/routes"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	e := echo.New()
	// e.Validator = validator.New()
	// pm := mPegawai.PegawaiModel{Db: db}
	um := mUser.New(db)
	uc := user.UserController{Repo: *um, Valid: validator.New()}

	routes.RegisterPath(e, uc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))

}
