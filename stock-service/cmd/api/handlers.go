package main

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/tonyproenca/stock-service/cmd/data"
	"net/http"
)

type JSONPostRequest struct {
	ProductName string `json:"productName"`
	ProductCode string `json:"productCode"`
	Quantity    int    `json:"quantity"`
}

func (app *Config) RetrieveStockProduct(w http.ResponseWriter, r *http.Request) {
	productCode := chi.URLParam(r, "productCode")

	stockProduct, err := app.Repo.GetOne(productCode)
	if err != nil {
		switch err.Error() {
		case "mongo: no documents in result":
			{
				_ = app.errorJSON(w, err, http.StatusNotFound)
				return
			}
		default:
			{
				_ = app.errorJSON(w, err, http.StatusInternalServerError)
				return
			}
		}
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Retrieving Product Stock",
		Data:    stockProduct,
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) StoreNewStockProduct(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPostRequest

	_ = app.readJSON(w, r, &requestPayload)
	stockProduct := data.StockProduct{
		ProductName: requestPayload.ProductName,
		ProductCode: requestPayload.ProductCode,
		Quantity:    requestPayload.Quantity,
	}

	err := app.Repo.Insert(stockProduct)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Product Stock data successfully stored",
	}

	_ = app.writeJSON(w, http.StatusCreated, resp)
}

func (app *Config) DeleteStockProduct(w http.ResponseWriter, r *http.Request) {
	productCode := chi.URLParam(r, "productCode")

	err := app.Repo.Delete(productCode)
	if err != nil {
		_ = app.errorJSON(w, errors.New("error deleting product code"), http.StatusInternalServerError)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Product Stock data successfully deleted",
	}

	_ = app.writeJSON(w, http.StatusOK, resp)
}

func (app *Config) UpdateStockProduct(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPostRequest

	_ = app.readJSON(w, r, &requestPayload)
	stockProduct := data.StockProduct{
		ProductName: requestPayload.ProductName,
		ProductCode: requestPayload.ProductCode,
		Quantity:    requestPayload.Quantity,
	}

	err := app.Repo.Update(stockProduct)

	if err != nil {
		switch err.Error() {
		case "not found":
			{
				_ = app.errorJSON(w, err, http.StatusNotFound)
				return
			}
		default:
			{
				_ = app.errorJSON(w, err, http.StatusInternalServerError)
				return
			}
		}
	}

	result := jsonResponse{
		Error:   false,
		Message: "Product Stock data successfully stored",
		Data:    stockProduct,
	}

	_ = app.writeJSON(w, http.StatusOK, result)
}
