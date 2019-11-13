package main

//https://jsonapi.org/format/1.0/
//https://github.com/HttpErrorPages/HttpErrorPages

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Change AVATAR  API
// @version 0.1
// @description 根据最近的节日生成特定的头像
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.github.com/go-boy
// @contact.email xinyong.wang@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /
func main() {
	hostPort := ":8080"
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
		HTML5:  false,
		// Index: "index.html",
	}))
	server := &http.Server{
		Addr:         hostPort,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	setupRouter(e)
	e.Logger.Fatal(e.StartServer(server))

}
