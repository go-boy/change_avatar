package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)
func main(){
	hostPort := 8080
	e := echo.New()
	e.HideBanner = true
	setupRouter(e)
	e.Logger.Info(" Listen and Server in ", hostPort)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", hostPort)))

}