package domain

import (
	"context"

	"github.com/datshiro/cyclo-ecommerce/internal/models"
)

type BrandRepo interface {
	CreateOne(ctx context.Context, brand *models.Brand) (*models.Brand, error)
	CreateMany(ctx context.Context, brands models.BrandSlice) error
}

type BrandUsecase interface {
	CreateOne(ctx context.Context, brand *models.Brand) (*models.Brand, error)
	CreateMany(ctx context.Context, brands models.BrandSlice) error
}
