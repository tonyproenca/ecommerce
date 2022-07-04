package main

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/tonyproenca/stock-service/cmd/data"
	"log"
	"net/http"
	"strings"
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
		if strings.Contains(err.Error(), "mongo: no documents in result") {
			_ = app.writeJSON(w, http.StatusNotFound, jsonResponse{
				Error:   true,
				Message: "Product not found",
			})
			return
		} else {
			_ = app.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

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
		if strings.Contains(err.Error(), "duplicate key error collection") {
			_ = app.writeJSON(w, http.StatusConflict, jsonResponse{
				Error:   true,
				Message: "There's an already existing product with the provided productCode",
			})
			return
		} else {
			_ = app.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	_ = app.writeJSON(w, http.StatusCreated, jsonResponse{
		Error:   false,
		Message: "Product Stock data successfully stored",
	})
}

func (app *Config) DeleteStockProduct(w http.ResponseWriter, r *http.Request) {
	productCode := chi.URLParam(r, "productCode")
	resp := jsonResponse{
		Error:   false,
		Message: "Product Stock data successfully deleted",
	}

	err := app.Repo.Delete(productCode)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			log.Println("No documents in result")
			_ = app.writeJSON(w, http.StatusOK, resp)
		}
		_ = app.errorJSON(w, errors.New("error deleting product code"), http.StatusInternalServerError)
		return
	}

	log.Println("Deleting product from database, productCode", productCode)
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
		if err.Error() == "mongo: no documents in result" {
			_ = app.writeJSON(w, http.StatusNotFound, jsonResponse{
				Error:   true,
				Message: "Product not found",
			})
			return
		} else {
			_ = app.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	_ = app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Product Stock data successfully stored",
		Data:    stockProduct,
	})
}
