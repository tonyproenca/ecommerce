package web

import (
	"testing"
)

func TestStoreNewStockProduct(t *testing.T) {
	//var serviceMock = new(service.ProductStockServiceMock)
	//handler := ProductStockHandler{serviceMock}
	//
	//postBody := map[string]interface{}{
	//	"ProductName": "test",
	//	"ProductCode": "123",
	//	"Quantity":    1,
	//}
	//
	//body, _ := json.Marshal(postBody)
	//req, _ := http.NewRequest("POST", "/stock-product", bytes.NewReader(body))
	//rr := httptest.NewRecorder()
	//handler = demo(http.HandlerFunc(handler.PostProductStock))
	//serviceMock.On("Insert", mock.Anything).Return(nil).Times(1)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusCreated, rr.Code)
	//assert.Contains(t, result, "Product Stock data successfully stored")
	//repositoryMock.AssertExpectations(t)
}

func TestStoreNewStockProductDatabaseErrorReturnsInternalServerError(t *testing.T) {
	//var testApp main.App
	//var repositoryMock = new(repository.MongoDBMock)
	//testApp.Repo = repositoryMock
	//
	//postBody := map[string]interface{}{
	//	"ProductName": "test",
	//	"ProductCode": "123",
	//	"Quantity":    1,
	//}
	//
	//body, _ := json.Marshal(postBody)
	//req, _ := http.NewRequest("POST", "/stock-product", bytes.NewReader(body))
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(testApp.StoreNewStockProduct)
	//
	//repositoryMock.On("Insert", mock.Anything).Return(errors.New("database error")).Times(1)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusInternalServerError, rr.Code)
	//assert.Equal(t, `{"error":true,"message":"database error"}`, result)
	//repositoryMock.AssertExpectations(t)
}

func TestRetrieveStockProduct(t *testing.T) {
	//var testApp main.App
	//var repositoryMock = new(repository.MongoDBMock)
	//testApp.Repo = repositoryMock
	//
	//req, _ := http.NewRequest("GET", "/stock-product/{productCode}", nil)
	//
	//rctx := chi.NewRouteContext()
	//rctx.URLParams.Add("productCode", "123")
	//
	//req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(testApp.RetrieveStockProduct)
	//
	//repositoryMock.On("GetOne", "123").Return(&repository.ProductStockEntity{
	//	ID:          "123",
	//	ProductCode: "1",
	//	ProductName: "product",
	//	Quantity:    0,
	//}, nil)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusOK, rr.Code)
	//assert.Contains(t, result, "Retrieving Product Stock")
	//assert.Contains(t, result, `"id":"123"`)
	//repositoryMock.AssertExpectations(t)
}

func TestRetrieveStockProductNotFound(t *testing.T) {
	//var testApp main.App
	//var repositoryMock = new(repository.MongoDBMock)
	//testApp.Repo = repositoryMock
	//
	//productCode := "123"
	//
	//req, _ := http.NewRequest("GET", "/stock-product/"+productCode, nil)
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(testApp.RetrieveStockProduct)
	//
	//repositoryMock.On("GetOne", mock.Anything).Return((*repository.ProductStockEntity)(nil), errors.New("database error")).Times(1)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusInternalServerError, rr.Code)
	//assert.Equal(t, `{"error":true,"message":"database error"}`, result)
	//repositoryMock.AssertExpectations(t)
}

func TestRetrieveStockProductDatabaseError(t *testing.T) {
	//var testApp main.App
	//var repositoryMock = new(repository.MongoDBMock)
	//testApp.Repo = repositoryMock
	//
	//productCode := "123"
	//
	//req, _ := http.NewRequest("GET", "/stock-product/"+productCode, nil)
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(testApp.RetrieveStockProduct)
	//
	//repositoryMock.On("GetOne", mock.Anything).Return((*repository.ProductStockEntity)(nil), errors.New("mongo: no documents in result")).Times(1)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusNotFound, rr.Code)
	//assert.Equal(t, `{"error":true,"message":"Product not found"}`, result)
	//repositoryMock.AssertExpectations(t)
}

func TestDeleteStockProduct(t *testing.T) {
	//var testApp main.App
	//var repositoryMock = new(repository.MongoDBMock)
	//testApp.Repo = repositoryMock
	//
	//productCode := "123"
	//
	//req, _ := http.NewRequest("DELETE", "/stock-product/"+productCode, nil)
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(testApp.DeleteStockProduct)
	//
	//repositoryMock.On("Delete", mock.Anything).Return(nil)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusOK, rr.Code)
	//assert.Equal(t, `{"error":false,"message":"Product Stock data successfully deleted"}`, result)
	//repositoryMock.AssertExpectations(t)
}

func TestDeleteStockProductDatabaseError(t *testing.T) {
	//var testApp main.App
	//var repositoryMock = new(repository.MongoDBMock)
	//testApp.Repo = repositoryMock
	//
	//productCode := "123"
	//
	//req, _ := http.NewRequest("DELETE", "/stock-product/"+productCode, nil)
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(testApp.DeleteStockProduct)
	//
	//repositoryMock.On("Delete", mock.Anything).Return(errors.New("database error")).Times(1)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusInternalServerError, rr.Code)
	//assert.Equal(t, `{"error":true,"message":"error deleting product code"}`, result)
	//repositoryMock.AssertExpectations(t)
}

func TestUpdateStockProduct(t *testing.T) {
	//var testApp main.App
	//var repositoryMock = new(repository.MongoDBMock)
	//testApp.Repo = repositoryMock
	//
	//postBody := map[string]interface{}{
	//	"ProductName": "test",
	//	"ProductCode": "123",
	//	"Quantity":    1,
	//}
	//
	//body, _ := json.Marshal(postBody)
	//req, _ := http.NewRequest("PUT", "/stock-product", bytes.NewReader(body))
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(testApp.UpdateStockProduct)
	//
	//repositoryMock.On("Update", mock.Anything).Return(nil).Times(1)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusOK, rr.Code)
	//assert.Contains(t, result, `Product Stock data successfully stored`)
	//repositoryMock.AssertExpectations(t)
}

func TestUpdateStockProductDatabaseError(t *testing.T) {
	//var testApp main.App
	//var repositoryMock = new(repository.MongoDBMock)
	//testApp.Repo = repositoryMock
	//
	//postBody := map[string]interface{}{
	//	"ProductName": "test",
	//	"ProductCode": "123",
	//	"Quantity":    1,
	//}
	//
	//body, _ := json.Marshal(postBody)
	//req, _ := http.NewRequest("PUT", "/stock-product", bytes.NewReader(body))
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(testApp.UpdateStockProduct)
	//
	//repositoryMock.On("Update", mock.Anything).Return(errors.New("database error")).Times(1)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusInternalServerError, rr.Code)
	//assert.Equal(t, `{"error":true,"message":"database error"}`, result)
	//repositoryMock.AssertExpectations(t)
}

func TestUpdateStockProductNotFound(t *testing.T) {
	//var testApp main.App
	//var repositoryMock = new(repository.MongoDBMock)
	//testApp.Repo = repositoryMock
	//
	//postBody := map[string]interface{}{
	//	"ProductName": "test",
	//	"ProductCode": "123",
	//	"Quantity":    1,
	//}
	//
	//body, _ := json.Marshal(postBody)
	//req, _ := http.NewRequest("PUT", "/stock-product", bytes.NewReader(body))
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(testApp.UpdateStockProduct)
	//
	//repositoryMock.On("Update", mock.Anything).Return(errors.New("mongo: no documents in result")).Times(1)
	//
	//handler.ServeHTTP(rr, req)
	//result := rr.Body.String()
	//
	//assert.Equal(t, http.StatusNotFound, rr.Code)
	//assert.Equal(t, `{"error":true,"message":"Product not found"}`, result)
	//repositoryMock.AssertExpectations(t)
}
