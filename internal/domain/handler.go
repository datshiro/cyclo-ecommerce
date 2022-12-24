package domain

import "github.com/labstack/echo/v4"

type Handler interface {
	Handle(echo.Context) error
	GetPath(prefixAPI string) string
}
