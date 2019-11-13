package handlers


import (
    "net/http"
    echo "github.com/labstack/echo/v4"
)

type httpError struct {
    code    int
    Key     string `json:"error"`
    Message string `json:"message"`
}

func newHTTPError(code int, key string, msg string) *httpError {
    return &httpError{
        code:    code,
        Key:     key,
        Message: msg,
    }
}

// Error 兼容 `error` 接口
func (e *httpError) Error() string {
    return e.Key + ": " + e.Message
}



// httpErrorHandler 定制 echo的HTTP错误handler.
func httpErrorHandler(err error, c echo.Context) {
    var (
        code = http.StatusInternalServerError
        key  = "ServerError"
        msg  string
    )

    if he, ok := err.(*httpError); ok {
        code = he.code
        key = he.Key
        msg = he.Message
    } else  {
		//debug
		msg = err.Error()

		// msg = http.StatusText(code)
    } 
        

    if !c.Response().Committed {
        if c.Request().Method == echo.HEAD {
            err := c.NoContent(code)
            if err != nil {
                c.Logger().Error(err)
            }
        } else {
            err := c.JSON(code, newHTTPError(code, key, msg))
            if err != nil {
                c.Logger().Error(err)
            }
        }
    }
}

