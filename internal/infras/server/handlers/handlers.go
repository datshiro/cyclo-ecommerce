package handlers

import (
	"github.com/datshiro/cyclo-ecommerce/internal/helper/db_helper.go"
	"github.com/datshiro/cyclo-ecommerce/internal/infras/server/config"
	"github.com/datshiro/cyclo-ecommerce/internal/infras/server/handlers/products"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, apiPrefix string, config config.Config) {
	database := db_helper.NewPGDB(config.DbUrl)

	products.RegisterHandlers(e, apiPrefix, database)
}
