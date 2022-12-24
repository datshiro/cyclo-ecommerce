package products

import (
	"database/sql"
	"log"

	"github.com/datshiro/cyclo-ecommerce/internal/domain"
	"github.com/datshiro/cyclo-ecommerce/internal/usecases"
	"github.com/datshiro/cyclo-ecommerce/internal/usecases/product/repo"
	"github.com/labstack/echo/v4"
)

const (
	ProductPath = "products"
)

func RegisterHandlers(e *echo.Echo, pathPrefix string, dbc *sql.DB) {

	repo := repo.New(dbc)

	uc := usecases.NewProductUsecase(repo)

	getHandlers := []domain.Handler{
		readHandler{uc: uc, path: ProductPath},
		readManyHandler{uc: uc, path: ProductPath},
	}

	// Register Get handlers
	for _, handler := range getHandlers {
		log.Println(handler.GetPath(pathPrefix))
		e.GET(handler.GetPath(pathPrefix), handler.Handle)
	}
}
