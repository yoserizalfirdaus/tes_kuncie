package repository

import "github.com/yoserizalfirdaus/tes_kuncie/entity"

type ProductRepository struct {
	sqldb SQLDatabase
}

type ProductRow struct {
	Sku      string  `db:"sku"`
	Name     string  `db:"name"`
	Price    float64 `db:"price"`
	Currency string  `db:"currency"`
}

func NewProductRepo(db SQLDatabase) (ProductRepository, error) {
	return ProductRepository{db}, nil
}

func (r ProductRepository) GetBySku(sku string) (entity.Product, error) {
	p := ProductRow{}
	err := r.sqldb.Get(&p, "SELECT sku, name, price, currency FROM product WHERE sku = $1", sku)
	if err != nil {
		return entity.Product{}, err
	}

	return p.toEntity(), err
}

func (r ProductRow) toEntity() entity.Product {
	return entity.Product{
		Sku:      r.Sku,
		Name:     r.Name,
		Price:    r.Price,
		Currency: r.Currency,
	}
}
