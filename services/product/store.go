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
        err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.CreatedAt)
        if err!= nil {
            return nil, err
        }
        products = append(products, p)

	}
	return products, nil

}

func (p *ProductStoreImpl) CreateProduct(prod *types.Product) (*types.Product, error) {
	stmt, err := p.db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES(?,?,?,?)")
    if err!= nil {
        return nil, err
    }
    res, err := stmt.Exec(&prod.Name, &prod.Description, &prod.Price, &prod.Quantity)
    if err!= nil {
        return nil, err
    }
    id, err := res.LastInsertId()
    if err!= nil {
        return nil, err
    }
    prod.ID = int(id)
    return prod, nil
}