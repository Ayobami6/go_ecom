package cart

import (
	"database/sql"

	"github.com/Ayobami6/go_ecom/types"
)

type CartStoreImpl struct {
	db *sql.DB
}

func NewCartStore(db *sql.DB) *CartStoreImpl {
    return &CartStoreImpl{db: db}
}

func (s *CartStoreImpl) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("Insert into orders (userId, total, status, address) values(?, ?, ?)", order.UserID, order.Total,order.Status, order.Address)
	if err!= nil {
        return 0, err
    }
	lastID, err := res.LastInsertId()
	if err!= nil {
        return 0, err
    }
	return int(lastID), nil
}

func (s *CartStoreImpl) CreateOrderItem(orderItem types.OrderItem) error {
	res, err := s.db.Exec("Insert into order_items (orderId, productId, quantity, price) values(?,?,?,?)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
    if err!= nil {
        return err
    }
    lastID, err := res.LastInsertId()
    if err!= nil {
        return err
    }
    orderItem.ID = int(lastID)
    return nil
}

