package repo

import (
	"context"
	"database/sql"

	"github.com/datshiro/cyclo-ecommerce/internal/domain"
	"github.com/datshiro/cyclo-ecommerce/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func New(dbc *sql.DB) domain.BrandRepo {
	return &branchRepo{dbc: dbc}
}

type branchRepo struct {
	dbc *sql.DB
}

func (b *branchRepo) CreateOne(ctx context.Context, branch *models.Brand) (*models.Brand, error) {
	if err := branch.Insert(ctx, b.dbc, boil.Infer()); err != nil {
		return nil, err
	}
	return branch, nil
}

func (b *branchRepo) CreateMany(ctx context.Context, branches models.BrandSlice) error {
	tx, err := b.dbc.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, branch := range branches {
		err := branch.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return tx.Rollback()
		}
	}
	return tx.Commit()

}
