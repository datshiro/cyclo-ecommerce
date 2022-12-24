package brand

import (
	"context"

	"github.com/datshiro/cyclo-ecommerce/internal/domain"
	"github.com/datshiro/cyclo-ecommerce/internal/models"
)

func New(repo domain.BrandRepo) domain.BrandUsecase {
	return &brandUC{repo: repo}
}

type brandUC struct {
	repo domain.BrandRepo
}

func (p *brandUC) CreateMany(ctx context.Context, brands models.BrandSlice) error {
	return p.repo.CreateMany(ctx, brands)
}

func (p *brandUC) CreateOne(ctx context.Context, brand *models.Brand) (*models.Brand, error) {
	return p.repo.CreateOne(ctx, brand)
}
