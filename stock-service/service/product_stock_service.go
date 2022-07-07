package service

import (
	"github.com/tonyproenca/stock-service/exceptions"
	"github.com/tonyproenca/stock-service/model"
	"github.com/tonyproenca/stock-service/repository"
	"log"
	"strings"
)

type ProductStockService struct {
	repo repository.Repository
}

type JsonResponse struct {
	Message string              `json:"message,omitempty"`
	Data    *model.ProductStock `json:"data,omitempty"`
}

func NewProductStockService(repository repository.Repository) *ProductStockService {
	return &ProductStockService{
		repo: repository,
	}
}

func (s *ProductStockService) RetrieveProductStock(productCode string) (*JsonResponse, error) {
	entity, err := s.repo.GetOne(productCode)

	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents in result") {
			log.Println("Fail to retrieve product. Product with code "+productCode+" not found.", err)
			return nil, &exceptions.NotFoundError{
				Detail: "No Product found for the product code " + productCode,
			}
		} else {
			log.Println("Fail to retrieve product", err)
			return nil, &exceptions.InternalServerError{
				Detail: "Something went wrong",
			}
		}
	}

	response := model.ProductStock{
		ProductCode: entity.ProductCode,
		ProductName: entity.ProductName,
		Quantity:    entity.Quantity,
	}

	return &JsonResponse{
		Message: "Retrieving Product Stock",
		Data:    &response,
	}, nil
}

func (s *ProductStockService) StoreNewProductStock(productStock model.ProductStock) (*JsonResponse, error) {
	err := s.repo.Insert(buildProductStockEntity(productStock))

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error collection") {
			log.Println("Fail to store new product. Product Code is already used.", err)
			return nil, &exceptions.ConflictError{
				Detail: "Product Code in use",
			}
		} else {
			log.Println("Failed to store new product.", err)
			return nil, &exceptions.InternalServerError{
				Detail: "Something went wrong",
			}
		}
	}

	return &JsonResponse{
		Message: "Product stock stored successfully",
		Data:    nil,
	}, nil
}

func buildProductStockEntity(productStock model.ProductStock) repository.ProductStockEntity {
	return repository.ProductStockEntity{
		ProductCode: productStock.ProductCode,
		ProductName: productStock.ProductName,
		Quantity:    productStock.Quantity,
	}
}

func (s *ProductStockService) DeleteProductStock(productCode string) (*JsonResponse, error) {
	err := s.repo.Delete(productCode)
	res := JsonResponse{
		Message: "Delete completed successfully",
		Data:    nil,
	}
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			log.Println("No documents in result. No record deleted")
			return &res, nil
		} else {
			log.Println("Failed to delete product.", err)
			return nil, &exceptions.InternalServerError{
				Detail: "Something went wrong",
			}
		}
	}

	return &res, nil
}

func (s *ProductStockService) UpdateProductStock(productStock model.ProductStock) (*JsonResponse, error) {
	err := s.repo.Update(buildProductStockEntity(productStock))
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents in result") {
			log.Println("Failed to update product. No product found for the product code "+productStock.ProductCode, err)
			return nil, &exceptions.NotFoundError{
				Detail: "No Product found for the product code " + productStock.ProductCode,
			}
		} else {
			log.Println("Failed to update product", err)
			return nil, &exceptions.InternalServerError{
				Detail: "Something went wrong",
			}
		}
	}

	return &JsonResponse{
		Message: "Product stock updated successfully",
		Data:    &productStock,
	}, nil
}
