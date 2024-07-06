package product

import (
	"database/sql"
	"fmt"
	"strings"

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


func (s *ProductStoreImpl) GetProductsByID(productIDs []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	// Convert productIDs to []interface{}
	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil

}

func (s *ProductStoreImpl) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = ?, price = ?, description = ?, quantity = ? WHERE id = ?", product.Name, product.Price,product.Description, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}