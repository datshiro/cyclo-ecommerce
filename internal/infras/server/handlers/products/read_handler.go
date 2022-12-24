package products

import (
	"net/http"
	"path"
	"strconv"

	"github.com/datshiro/cyclo-ecommerce/internal/consts"
	"github.com/datshiro/cyclo-ecommerce/internal/domain"
	"github.com/labstack/echo/v4"
)

///////////////////// READ ONE HANDLER
type readHandler struct {
	path string
	uc   domain.ProductUsecase
}

func (h readHandler) GetPath(prefixAPI string) string {
	return path.Join(prefixAPI, h.path, ":id")
}

func (h readHandler) Handle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if id == 0 {
		return consts.ErrInvalidRequest
	}

	product, err := h.uc.GetOneById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, WriteResponse(product))
}

///////////////////// READ MANY HANDLER
type readManyHandler struct {
	uc   domain.ProductUsecase
	path string
}

func (h readManyHandler) GetPath(prefixAPI string) string {
	return path.Join(prefixAPI, h.path)
}

func (h readManyHandler) Handle(c echo.Context) error {
	request := NewRequest()

	if err := request.Bind(c); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return err
	}

	products, err := h.uc.GetManyWithFilters(
		c.Request().Context(),
		request.GetBrandIDs(),
		request.GetMinPrice(),
		request.GetMaxPrice(),
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, WriteManyProductResponse(products, request.GetSortBy()))
}
