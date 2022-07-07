package repository

import (
	"github.com/stretchr/testify/mock"
)

type MongoDBMock struct {
	mock.Mock
}

func (_m *MongoDBMock) Insert(stockProduct ProductStockEntity) error {
	args := _m.Called(stockProduct)
	return args.Error(0)
}

func (_m *MongoDBMock) GetOne(productCode string) (*ProductStockEntity, error) {
	args := _m.Called(productCode)
	return args.Get(0).(*ProductStockEntity), args.Error(1)
}

func (_m *MongoDBMock) Update(stockProduct ProductStockEntity) error {
	args := _m.Called(stockProduct)
	return args.Error(0)
}

func (_m *MongoDBMock) Delete(productCode string) error {
	args := _m.Called(productCode)
	return args.Error(0)
}
