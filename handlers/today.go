package handlers

import (
	"time"
	"net/http"
	echo "github.com/labstack/echo/v4"
	sl "github.com/go-boy/change_avatar/libs/solarlunar"
)

// TodayHandler get today info
func TodayHandler(e echo.Context) error{
	var timeLocation *time.Location
	timeLocation, _ = time.LoadLocation("Asia/Shanghai")
	var today time.Time
	today = time.Now()
	today.In(timeLocation)
	todayStr :=  today.Format("2006-01-02")
	lunar := sl.SolarToChineseLuanr(todayStr)
	lunarSimple := sl.SolarToSimpleLuanr(todayStr)
	msg := map[string]interface{}{"today": todayStr,"lunar":lunar,"simple_lunar":lunarSimple}
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": msg})
}
