package products

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/datshiro/cyclo-ecommerce/internal/mock/mock_domain"
	"github.com/datshiro/cyclo-ecommerce/internal/models"
	"github.com/datshiro/cyclo-ecommerce/internal/usecases"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockProduct = &models.Product{
		ID:    1,
		Name:  "Product 1",
		Price: 1,
	}

	productJSON = `{"id":1,"name":"Product 1","price":1,"brand":""}
`
)

func TestReadOneSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := mock_domain.NewMockProductRepo(mockCtl)
	mock.EXPECT().GetOneById(context.Background(), gomock.Eq(1)).Return(mockProduct, nil)

	uc := usecases.NewProductUsecase(mock)
	h := &readHandler{path: ProductPath, uc: uc}

	// Assertions
	if assert.NoError(t, h.Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, productJSON, rec.Body.String())
	}
}
