package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetAllUsers(c echo.Context) error {
	fmt.Println("url ", c.Request().URL)
	return c.JSON(http.StatusOK, "okokok")
}

func (h *Handler) GetUser(c echo.Context) error {
	return c.JSON(http.StatusAccepted, "hello")
}
