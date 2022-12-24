package product

import (
	"context"

	"github.com/datshiro/cyclo-ecommerce/internal/domain"
	"github.com/datshiro/cyclo-ecommerce/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func New(repo domain.ProductRepo) domain.ProductUsecase {
	return &productUC{repo: repo}
}

type productUC struct {
	repo domain.ProductRepo
}

func (p *productUC) CreateMany(ctx context.Context, products models.ProductSlice) error {
	return p.repo.CreateMany(ctx, products)
}

func (p *productUC) CreateOne(ctx context.Context, product *models.Product) (*models.Product, error) {
	return p.repo.CreateOne(ctx, product)
}

func (p *productUC) GetOneById(ctx context.Context, id int) (*models.Product, error) {
	return p.repo.GetOneById(ctx, id)
}

func (p *productUC) GetMany(ctx context.Context, mods ...qm.QueryMod) (models.ProductSlice, error) {
	return p.repo.GetMany(ctx, mods...)
}

func (p *productUC) GetManyWithFilters(ctx context.Context, brandIDs []int, minPrice, maxPrice float64) (models.ProductSlice, error) {
	mods := []qm.QueryMod{}
	if len(brandIDs) > 0 {
		mods = append(mods, models.ProductWhere.BrandID.IN(brandIDs))
	}
	if minPrice != 0 {
		mods = append(mods, models.ProductWhere.Price.GTE(minPrice))
	}
	if maxPrice != 0 {
		mods = append(mods, models.ProductWhere.Price.LTE(maxPrice))
	}
	return p.repo.GetMany(ctx, mods...)
}
