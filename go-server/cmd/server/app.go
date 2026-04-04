package main

import "github.com/labstack/echo/v4"

type APP struct {
	echo *echo.Echo
}

func NewApp() *APP {
	app := echo.New()

	return &APP{
		echo: app,
	}
}
