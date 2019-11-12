package handlers

import (
	"net/http"
	echo "github.com/labstack/echo/v4"
)

// RecentFestivals get recent festivals
func RecentFestivals(e echo.Context) error{
	today := GetBJToday()
	todayStr :=  today.Format("2006-01-02")
	msg := map[string]interface{}{"today": todayStr}
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": msg})
}
