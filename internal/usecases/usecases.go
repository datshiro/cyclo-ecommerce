package usecases

import (
	"github.com/datshiro/cyclo-ecommerce/internal/usecases/brand"
	"github.com/datshiro/cyclo-ecommerce/internal/usecases/product"
)

var (
	NewProductUsecase = product.New
	NewBrandUsecase   = brand.New
)
