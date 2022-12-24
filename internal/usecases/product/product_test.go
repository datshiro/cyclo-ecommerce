package product

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/datshiro/cyclo-ecommerce/internal/domain"
	"github.com/datshiro/cyclo-ecommerce/internal/mock/mock_domain"
	"github.com/datshiro/cyclo-ecommerce/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
	args []int
}
type TestSuite struct {
	suite.Suite
	uc   domain.ProductUsecase
	repo *mock_domain.MockProductRepo
}

var (
	mockProduct = &models.Product{
		ID:        1,
		Name:      "Mock Name",
		Price:     123,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: null.TimeFrom(time.Time{}),
	}
)

func (s *TestSuite) SetupTest() {

	mockCtl := gomock.NewController(s.T())

	repo := mock_domain.NewMockProductRepo(mockCtl)
	uc := New(repo)
	defer mockCtl.Finish()

	s.uc = uc
	s.repo = repo
}

func (s *TestSuite) TestGetOne() {

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				s.repo.EXPECT().GetOneById(context.Background(), 1).Return(new(models.Product), nil)
			},
			res:  new(models.Product),
			err:  nil,
			args: []int{1},
		},
		{
			name: "return result",
			mock: func() {
				s.repo.EXPECT().GetOneById(context.Background(), 1).Return(&models.Product{ID: 1, Name: "Product 1", Price: 1234}, nil)
			},
			res:  &models.Product{ID: 1, Name: "Product 1", Price: 1234},
			err:  nil,
			args: []int{1},
		},
		{
			name: "result with error",
			mock: func() {
				s.repo.EXPECT().GetOneById(context.Background(), 0).Return(new(models.Product), errInternalServErr)
			},
			res:  new(models.Product),
			err:  errInternalServErr,
			args: []int{0},
		},
	}

	for _, tc := range tests {
		tc := tc
		fmt.Println("name", tc.name)

		tc.mock()

		res, err := s.uc.GetOneById(context.Background(), tc.args[0])

		t := s.T()
		require.Equal(t, tc.res, res)
		require.ErrorIs(t, err, tc.err)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestProductUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
