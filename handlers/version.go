package handlers

import (
	"net/http"
	echo "github.com/labstack/echo/v4"
)

const version = "0.1"
const releaseNote = "test"
// Version get current version
func Version(e echo.Context) error{
	msg := map[string]interface{}{"version": version, "release_note": releaseNote}
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": msg})
}
