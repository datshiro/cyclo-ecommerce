package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/datshiro/cyclo-ecommerce/internal/helper/db_helper.go"
	"github.com/datshiro/cyclo-ecommerce/internal/infras/server/config"
	"github.com/datshiro/cyclo-ecommerce/internal/models"
	"github.com/datshiro/cyclo-ecommerce/internal/usecases"
	brand_repo "github.com/datshiro/cyclo-ecommerce/internal/usecases/brand/repo"
	"github.com/datshiro/cyclo-ecommerce/internal/usecases/product/repo"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	pflag.String("env", "staging", "Env mode")
	pflag.String("configPath", "./", "Env directory path")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	config, err := config.NewConfig(
		viper.GetString("configPath"),
		viper.GetString("env"),
	)
	if err != nil {
		log.Fatalf("Failed to init configuration: %v", err)
	}

	log.Printf("Config %v", config)
	ctx := context.Background()
	dbc := db_helper.NewPGDB(config.DbUrl)
	if err := SeedBrands(ctx, dbc); err != nil {
		log.Fatalf("Failed to SeedProducts: %v", err)
	}
	if err := SeedProducts(ctx, dbc); err != nil {
		log.Fatalf("Failed to SeedProducts: %v", err)
	}

}

var (
	BrandList = []string{"Zara", "Prada", "Pedro", "H&M", "P&B"}
)

func SeedBrands(ctx context.Context, dbc *sql.DB) error {
	brandRepo := brand_repo.New(dbc)
	uc := usecases.NewBrandUsecase(brandRepo)
	brands := models.BrandSlice{}
	for _, brandName := range BrandList {
		brands = append(brands, &models.Brand{Name: brandName})
	}
	if err := uc.CreateMany(ctx, brands); err != nil {
		return err
	}
	log.Printf("Inserted %v rows of brands", len(brands))
	return nil
}

func SeedProducts(ctx context.Context, dbc *sql.DB) error {
	productRepo := repo.New(dbc)
	uc := usecases.NewProductUsecase(productRepo)

	products := models.ProductSlice{}
	for i := 1; i <= 100; i++ {
		brandId := (i % (len(BrandList) - 1)) + 1
		fmt.Println(brandId)
		products = append(products, &models.Product{
			ID:      i,
			Name:    fmt.Sprintf("Product %d", i),
			Price:   1.0 * float64(i),
			BrandID: brandId,
		})
	}
	if err := uc.CreateMany(ctx, products); err != nil {
		return err
	}
	log.Printf("Inserted %v rows of products", len(products))
	return nil
}
