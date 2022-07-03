package data

import (
	"github.com/stretchr/testify/mock"
)

type MongoDBMock struct {
	mock.Mock
}

func (_m *MongoDBMock) Insert(stockProduct StockProduct) error {
	args := _m.Called(stockProduct)
	return args.Error(0)
}

func (_m *MongoDBMock) GetOne(productCode string) (*StockProduct, error) {
	args := _m.Called(productCode)
	return args.Get(0).(*StockProduct), args.Error(1)
}

func (_m *MongoDBMock) Update(stockProduct StockProduct) error {
	args := _m.Called(stockProduct)
	return args.Error(0)
}

func (_m *MongoDBMock) Delete(productCode string) error {
	args := _m.Called(productCode)
	return args.Error(0)
}
