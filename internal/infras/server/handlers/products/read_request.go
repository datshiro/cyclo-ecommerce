package products

import (
	"github.com/labstack/echo/v4"
)

func NewRequest() Request {
	return &request{}
}

type Request interface {
	Bind(ctx echo.Context) error
	Validate() error
	GetMinPrice() float64
	GetMaxPrice() float64
	GetBrandIDs() []int
	GetSortBy() string
}

type request struct {
	BrandIds []int   `json:"brand_ids" param:"brand_ids" query:"brand_ids"`
	MinPrice float64 `json:"min_price" param:"min_price" query:"min_price"`
	MaxPrice float64 `json:"max_price" param:"max_price" query:"max_price"`
	SortBy   string  `json:"sort_by" param:"sort_by" query:"sort_by"`
}

func (r *request) GetMinPrice() float64 {
	return r.MinPrice
}

func (r *request) GetMaxPrice() float64 {
	return r.MaxPrice
}

func (r *request) GetBrandIDs() []int {
	return r.BrandIds
}

func (r *request) Bind(ctx echo.Context) error {
	return ctx.Bind(r)
}

func (r *request) GetSortBy() string {
	return r.SortBy
}
func (r *request) Validate() error {
	return nil
}
