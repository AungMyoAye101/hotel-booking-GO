package response

import "github.com/labstack/echo/v4"

type StandardResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Result     any    `json:"result"`
	Meta       any    `json:"meta,omitempty"`
}

func JSON(
	c echo.Context,
	statusCode int,
	success bool,
	message string,
	result any,
	meta any,
) error {
	return c.JSON(statusCode, StandardResponse{
		Success:    success,
		Message:    message,
		StatusCode: statusCode,
		Result:     result,
		Meta:       meta,
	})
}

func SuccessResponse(c echo.Context, statusCode int, message string, data any) error {
	return JSON(c, statusCode, true, message, data, nil)
}

func SuccessResponseWithMeta(c echo.Context, statusCode int, message string, data any, meta any) error {
	return JSON(c, statusCode, true, message, data, meta)
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return JSON(c, statusCode, false, message, nil, nil)
}
