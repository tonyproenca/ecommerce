package repository

type Repository interface {
	Insert(productStock ProductStockEntity) error
	GetOne(productCode string) (*ProductStockEntity, error)
	Update(productStock ProductStockEntity) error
	Delete(productCode string) error
}
