package service

import (
	"github.com/tonyproenca/stock-service/model"
)

type Service interface {
	RetrieveProductStock(productCode string) (*JsonResponse, error)
	StoreNewProductStock(productStock model.ProductStock) (*JsonResponse, error)
	DeleteProductStock(productCode string) (*JsonResponse, error)
	UpdateProductStock(productStock model.ProductStock) (*JsonResponse, error)
}
