package handlers

import (
	"net/http"
	echo "github.com/labstack/echo/v4"
)

const version = "0.1"
const releaseNote = "test"

// Version get current version
// @Summary 获取发版信息
// @Description get 获取当前版本及发版信息
// @ID get-current-version-
// @Accept  json
// @Produce  json
// @Router /api/v1.0/version [get]
func Version(e echo.Context) error{
	msg := map[string]interface{}{"version": version, "release_note": releaseNote}
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": msg})
}
