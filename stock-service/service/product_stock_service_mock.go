package service

import (
	"github.com/stretchr/testify/mock"
	"github.com/tonyproenca/stock-service/model"
)

type ProductStockServiceMock struct {
	mock.Mock
}

func (s *ProductStockServiceMock) RetrieveProductStock(productCode string) (*JsonResponse, error) {
	args := s.Called(productCode)
	return args.Get(0).(*JsonResponse), args.Error(0)
}

func (s *ProductStockServiceMock) StoreNewProductStock(productStock model.ProductStock) (*JsonResponse, error) {
	args := s.Called(productStock)
	return args.Get(0).(*JsonResponse), args.Error(0)
}

func (s *ProductStockServiceMock) UpdateProductStock(productStock model.ProductStock) (*JsonResponse, error) {
	args := s.Called(productStock)
	return args.Get(0).(*JsonResponse), args.Error(0)
}

func (s *ProductStockServiceMock) DeleteProductStock(productCode string) (*JsonResponse, error) {
	args := s.Called(productCode)
	return args.Get(0).(*JsonResponse), args.Error(0)
}
