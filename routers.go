package main

import (
	"net/http"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/go-boy/change_avatar/handlers"
)


// setupRouter 初始化路由
func setupRouter(e *echo.Echo) {
	// Ping test
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	api1_0 := e.Group("/api/v1.0")
	api1_0.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "tacey" && password == "wong" {
			return true, nil
		}
		return false, nil
	}))
	{
		api1_0.GET("/version", handlers.Version)
		api1_0.GET("/today",handlers.TodayHandler)
		api1_0.GET("/recentfestivals",handlers.RecentFestivals)


	}
}


