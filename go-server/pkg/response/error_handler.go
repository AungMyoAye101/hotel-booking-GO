package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	statusCode := http.StatusInternalServerError
	message := http.StatusText(statusCode)

	if he, ok := err.(*echo.HTTPError); ok {
		statusCode = he.Code
		switch v := he.Message.(type) {
		case string:
			message = v
		case error:
			message = v.Error()
		default:
			message = fmt.Sprint(v)
		}
	} else if err != nil {
		message = err.Error()
	}

	_ = ErrorResponse(c, statusCode, message)
}
