package types

import "time"

type UserStore interface {
	CreateUser(user *User) error
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)

}

type ProductStore interface {
	GetProducts() ([]*Product, error)
	CreateProduct(product *Product) (*Product, error)
	GetProductsByID(ps []int)([]Product, error)
	UpdateProduct(Product) error
}

type CartStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}


type Order struct {
	ID          int    `json:"id"`
	UserID       int    `json:"userId"`
	Total   	float64 `json:"total"`
	Status       string `json:"status"`
	Address string `json:"address"`
	CreatedAt time.Time `json:"createdAt"`

}

type OrderItem struct {
	ID          int    `json:"id"`
	OrderID     int    `json:"orderId"`
	ProductID    int    `json:"productId"`
	Quantity    int    `json:"quantity"`
	Price       float64 `json:"price"`
	CreatedAt time.Time `json:"createdAt"`

}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int    `json:"quantity"`
	// Image    string `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
}
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt time.Time`json:"createdAt"`
}

type RegisterUserPayLoad struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=20"`
}

type LoginUserPayLoad struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type ProductCreatePayLoad struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	// Image    string `json:"image"`
}

type CartCheckoutItem struct {
	ProductID int `json:"productId"`
	Quantity int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items []CartCheckoutItem `json:"items"`
}