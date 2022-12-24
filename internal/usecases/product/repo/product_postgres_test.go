package repo

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/datshiro/cyclo-ecommerce/internal/models"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
)

type TestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	db   *sql.DB
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

	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	s.NoError(err)
	s.mock = mock
	s.db = db
}

func (s *TestSuite) TestCreateOneSuccess() {
	// creat mock data for test
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(mockProduct.ID, mockProduct.Name, mockProduct.Price)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM cyclo.products").
		WithArgs(mockProduct.ID).
		WillReturnRows(rows)

	s.mock.ExpectCommit()

	repo := New(s.db)
	// execute test
	_, err := repo.GetOneById(context.Background(), mockProduct.ID)
	if err != nil {
		fmt.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func (s *TestSuite) TestCreateOneFailed() {

	// creat mock data for test
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO cyclo.products").
		WithArgs(mockProduct.Name, mockProduct.Price).
		WillReturnError(fmt.Errorf("some error"))
	s.mock.ExpectRollback()

	s.mock.ExpectCommit()

	repo := New(s.db)
	// execute test
	_, err := repo.CreateOne(context.Background(), &models.Product{Name: mockProduct.Name, Price: mockProduct.Price})
	if err != nil {
		fmt.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (s *TestSuite) TestGetOneSuccess() {
	// creat mock data for test
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO cyclo.products").
		WithArgs(mockProduct.Name, mockProduct.Price).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)
	s.mock.ExpectCommit()

	repo := New(s.db)
	// execute test
	_, err := repo.CreateOne(context.Background(), &models.Product{Name: mockProduct.Name, Price: mockProduct.Price})
	if err != nil {
		fmt.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (s *TestSuite) TestGetOneFailed() {
	// creat mock data for test

	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM cyclo.products").
		WillReturnError(fmt.Errorf("some error"))

	s.mock.ExpectCommit()

	repo := New(s.db)
	// execute test
	_, err := repo.GetMany(context.Background())
	if err != nil {
		fmt.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func (s *TestSuite) TestGetManyOneSuccess() {
	// creat mock data for test
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(mockProduct.ID, mockProduct.Name, mockProduct.Price).
		AddRow(2, "name2", 456)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM cyclo.products").
		WillReturnRows(rows)

	s.mock.ExpectCommit()

	repo := New(s.db)
	// execute test
	products, err := repo.GetMany(context.Background())
	if err != nil {
		fmt.Errorf("error was not expected while updating stats: %s", err)
	}
	fmt.Println(products)

	// we make sure that all expectations were met
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Errorf("there were unfulfilled expectations: %s", err)
	}

}

// func (s *TestSuite) TestGetManyOneFailed() {
// }

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestProductRepoSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
