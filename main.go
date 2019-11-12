package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

	_ "github.com/go-boy/change_avatar/docs"
)
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main(){
	hostPort := 8080
	e := echo.New()
	e.HideBanner = true
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	setupRouter(e)
	e.Logger.Info(" Listen and Server in ", hostPort)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", hostPort)))

}