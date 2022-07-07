package web

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/tonyproenca/stock-service/exceptions"
	"github.com/tonyproenca/stock-service/model"
	"github.com/tonyproenca/stock-service/service"
	"net/http"
)

type ProductStockHandler struct {
	productStockService service.Service
}

type JSONPostRequest struct {
	ProductName string `json:"productName"`
	ProductCode string `json:"productCode"`
	Quantity    int    `json:"quantity"`
}

func NewProductStockHandler(srv service.Service) *ProductStockHandler {
	return &ProductStockHandler{
		productStockService: srv,
	}
}

func (h *ProductStockHandler) GetProductStock(w http.ResponseWriter, r *http.Request) {
	productCode := chi.URLParam(r, "productCode")

	response, err := h.productStockService.RetrieveProductStock(productCode)
	if err != nil {
		errorsHandler(h, w, err)
		return
	}

	_ = h.writeJSON(w, http.StatusOK, response)
}

func (h *ProductStockHandler) PostProductStock(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPostRequest
	_ = h.readJSON(w, r, &requestPayload)

	productStock := buildProductStockFromJSON(requestPayload)

	res, err := h.productStockService.StoreNewProductStock(productStock)
	if err != nil {
		errorsHandler(h, w, err)
		return
	}

	_ = h.writeJSON(w, http.StatusCreated, res)
}

func buildProductStockFromJSON(requestPayload JSONPostRequest) model.ProductStock {
	return model.ProductStock{
		ProductName: requestPayload.ProductName,
		ProductCode: requestPayload.ProductCode,
		Quantity:    requestPayload.Quantity,
	}
}

func (h *ProductStockHandler) UpdateProductStock(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPostRequest
	_ = h.readJSON(w, r, &requestPayload)

	productStock := buildProductStockFromJSON(requestPayload)

	response, err := h.productStockService.UpdateProductStock(productStock)
	if err != nil {
		errorsHandler(h, w, err)
		return
	}

	_ = h.writeJSON(w, http.StatusOK, response)
}

func (h *ProductStockHandler) DeleteProductStock(w http.ResponseWriter, r *http.Request) {
	productCode := chi.URLParam(r, "productCode")

	res, err := h.productStockService.DeleteProductStock(productCode)
	if err != nil {
		errorsHandler(h, w, err)
		return
	}

	_ = h.writeJSON(w, http.StatusOK, res)
}

func errorsHandler(h *ProductStockHandler, w http.ResponseWriter, err error) {
	var statusNotFoundErrorPtr *exceptions.NotFoundError
	var statusConflictErrorPtr *exceptions.ConflictError

	if errors.As(err, &statusNotFoundErrorPtr) {
		_ = h.errorJSON(w, err, http.StatusNotFound)
	} else if errors.As(err, &statusConflictErrorPtr) {
		_ = h.errorJSON(w, err, http.StatusConflict)
	} else {
		_ = h.errorJSON(w, err, http.StatusInternalServerError)
	}
}
