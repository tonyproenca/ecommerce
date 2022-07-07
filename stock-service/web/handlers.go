package web

import "net/http"

type Handlers interface {
	GetProductStock(w http.ResponseWriter, r *http.Request)
	PostProductStock(w http.ResponseWriter, r *http.Request)
	UpdateProductStock(w http.ResponseWriter, r *http.Request)
	DeleteProductStock(w http.ResponseWriter, r *http.Request)
}
