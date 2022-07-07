package model

type ProductStock struct {
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
}
