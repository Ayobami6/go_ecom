package product

import (
	"database/sql"

	"github.com/Ayobami6/go_ecom/types"
)

type ProductStoreImpl struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStoreImpl {
    return &ProductStoreImpl{db: db}
}

// interface impls

func (p *ProductStoreImpl) GetProducts() ([]*types.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products")
	if err!= nil {
        return nil, err
    }
	products := make([]*types.Product, 0)
	for rows.Next() {
		var p = new(types.Product)
        err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Image, &p.CreatedAt)
        if err!= nil {
            return nil, err
        }
        products = append(products, p)

	}
	return products, nil

}