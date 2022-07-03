package data

type Repository interface {
	Insert(stockProduct StockProduct) error
	GetOne(productCode string) (*StockProduct, error)
	Update(stockProduct StockProduct) error
	Delete(productCode string) error
}
