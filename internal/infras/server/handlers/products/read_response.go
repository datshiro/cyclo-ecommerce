package products

import (
	"sort"
	"strings"

	"github.com/datshiro/cyclo-ecommerce/internal/models"
)

type ProductResponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price" `
	Brand string  `json:"brand" omitempty:"brand"`
}

func WriteResponse(product *models.Product) ProductResponse {
	response := ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
	if product.R != nil {
		response.Brand = product.R.Brand.Name
	}
	return response
}

// Response many products
type ManyProductReponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func WriteManyProductResponse(products models.ProductSlice, sortBy string) []ManyProductReponse {
	sort.SliceStable(products, func(i, j int) bool {
		switch strings.ToLower(sortBy) {
		case "brands":
			return products[i].BrandID < products[j].BrandID
		case "price":
			return products[i].Price < products[j].Price
		default:
			return products[i].Price < products[j].Price
		}
	})
	response := []ManyProductReponse{}
	for _, product := range products {
		response = append(response,
			ManyProductReponse{
				ID:    product.ID,
				Name:  product.Name,
				Price: product.Price,
			})
	}
	return response
}
