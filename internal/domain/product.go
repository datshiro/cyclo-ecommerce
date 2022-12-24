package domain

import (
	"context"

	"github.com/datshiro/cyclo-ecommerce/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ProductRepo interface {
	CreateOne(ctx context.Context, product *models.Product) (*models.Product, error)
	CreateMany(ctx context.Context, products models.ProductSlice) error
	GetOneById(ctx context.Context, id int) (*models.Product, error)
	GetMany(ctx context.Context, mods ...qm.QueryMod) (models.ProductSlice, error)
}

type ProductUsecase interface {
	CreateOne(ctx context.Context, product *models.Product) (*models.Product, error)
	CreateMany(ctx context.Context, products models.ProductSlice) error
	GetOneById(ctx context.Context, id int) (*models.Product, error)
	GetMany(ctx context.Context, mods ...qm.QueryMod) (models.ProductSlice, error)
	GetManyWithFilters(ctx context.Context, brandIDs []int, minPrice, maxPrice float64) (models.ProductSlice, error)
}
