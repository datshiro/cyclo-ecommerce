package repo

import (
	"context"
	"database/sql"

	"github.com/datshiro/cyclo-ecommerce/internal/domain"
	"github.com/datshiro/cyclo-ecommerce/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func New(dbc *sql.DB) domain.ProductRepo {
	return &productRepo{dbc: dbc}
}

type productRepo struct {
	dbc *sql.DB
}

func (p *productRepo) CreateMany(ctx context.Context, products models.ProductSlice) error {
	tx, err := p.dbc.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	for _, product := range products {
		err := product.Insert(ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (p *productRepo) CreateOne(ctx context.Context, product *models.Product) (*models.Product, error) {
	if err := product.Insert(ctx, p.dbc, boil.Infer()); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepo) GetOneById(ctx context.Context, id int) (*models.Product, error) {
	return models.Products(
		models.ProductWhere.ID.EQ(id),
		qm.Load("Brand"),
	).One(ctx, p.dbc)
}

func (p *productRepo) GetMany(ctx context.Context, mods ...qm.QueryMod) (models.ProductSlice, error) {
	return models.Products(mods...).All(ctx, p.dbc)
}
