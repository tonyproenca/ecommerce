package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tonyproenca/stock-service/cmd/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoreNewStockProduct(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	postBody := map[string]interface{}{
		"ProductName": "test",
		"ProductCode": "123",
		"Quantity":    "1",
	}

	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("POST", "/stock-product", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.StoreNewStockProduct)

	repositoryMock.On("Insert", mock.Anything).Return(nil).Times(1)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Response code should be StatusCreated")
	// TODO assert response
}

//func TestStoreNewStockProductMissingPropertiesOnRequestBodyRetrievesBadRequest(t *testing.T) {
//	var testApp Config
//	var repositoryMock = new(data.MongoDBMock)
//	testApp.Repo = repositoryMock
//
//	postBody := map[string]interface{}{
//		"ProductName": "test",
//		"Quantity":    "1",
//	}
//
//	body, _ := json.Marshal(postBody)
//	req, _ := http.NewRequest("POST", "/stock-product", bytes.NewReader(body))
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(testApp.StoreNewStockProduct)
//
//	handler.ServeHTTP(rr, req)
//
//	assert.Equal(t, http.StatusBadRequest, rr.Code, "Request Body should contain ProductName, ProductCode and Quantity properties")
//	// TODO assert response
//	repositoryMock.AssertExpectations(t)
//}

func TestStoreNewStockProductDatabaseErrorReturnsInternalServerError(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	postBody := map[string]interface{}{
		"ProductName": "test",
		"ProductCode": "123",
		"Quantity":    "1",
	}

	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("POST", "/stock-product", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.StoreNewStockProduct)

	repositoryMock.On("Insert", mock.Anything).Return(errors.New("database Error")).Times(1)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	// TODO assert response
	repositoryMock.AssertExpectations(t)
}

func TestRetrieveStockProduct(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	productCode := "123"

	req, _ := http.NewRequest("GET", "/stock-product/"+productCode, nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.RetrieveStockProduct)

	repositoryMock.On("GetOne", mock.Anything).Return(&data.StockProduct{
		ID:          "123",
		ProductCode: "1",
		ProductName: "product",
		Quantity:    0,
	}, nil)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// TODO assert response
	repositoryMock.AssertExpectations(t)
}

func TestRetrieveStockProductReturningNotFound(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	productCode := "123"

	req, _ := http.NewRequest("GET", "/stock-product/"+productCode, nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.RetrieveStockProduct)

	repositoryMock.On("GetOne", mock.Anything).Return((*data.StockProduct)(nil), errors.New("not found")).Times(1)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	// TODO assert response
	repositoryMock.AssertExpectations(t)
}

func TestRetrieveStockProductDatabaseError(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	productCode := "123"

	req, _ := http.NewRequest("GET", "/stock-product/"+productCode, nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.RetrieveStockProduct)

	repositoryMock.On("GetOne", mock.Anything).Return((*data.StockProduct)(nil), errors.New("database error")).Times(1)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	// TODO assert response
	repositoryMock.AssertExpectations(t)
}

func TestDeleteStockProduct(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	productCode := "123"

	req, _ := http.NewRequest("DELETE", "/stock-product/"+productCode, nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.DeleteStockProduct)

	repositoryMock.On("Delete", mock.Anything).Return(nil)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// TODO assert response
	repositoryMock.AssertExpectations(t)
}

func TestDeleteStockProductDatabaseError(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	productCode := "123"

	req, _ := http.NewRequest("DELETE", "/stock-product/"+productCode, nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.DeleteStockProduct)

	repositoryMock.On("Delete", mock.Anything).Return(errors.New("database error")).Times(1)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	// TODO assert response
	repositoryMock.AssertExpectations(t)
}

func TestUpdateStockProduct(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	postBody := map[string]interface{}{
		"ProductName": "test",
		"ProductCode": "123",
		"Quantity":    "1",
	}

	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("PUT", "/stock-product", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.UpdateStockProduct)

	repositoryMock.On("Update", mock.Anything).Return(nil).Times(1)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// TODO assert response
	repositoryMock.AssertExpectations(t)
}

func TestUpdateStockProductDatabaseError(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	postBody := map[string]interface{}{
		"ProductName": "test",
		"ProductCode": "123",
		"Quantity":    "1",
	}

	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("PUT", "/stock-product", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.UpdateStockProduct)

	repositoryMock.On("Update", mock.Anything).Return(errors.New("database error")).Times(1)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	// TODO assert response
	repositoryMock.AssertExpectations(t)
}

func TestUpdateStockProductNotFound(t *testing.T) {
	var testApp Config
	var repositoryMock = new(data.MongoDBMock)
	testApp.Repo = repositoryMock

	postBody := map[string]interface{}{
		"ProductName": "test",
		"ProductCode": "123",
		"Quantity":    "1",
	}

	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("PUT", "/stock-product", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.UpdateStockProduct)

	repositoryMock.On("Update", mock.Anything).Return(errors.New("not found")).Times(1)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	// TODO assert response
	repositoryMock.AssertExpectations(t)
}
