package types

import "time"

type UserStore interface {
	CreateUser(user *User) error
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)

}

type ProductStore interface {
	GetProducts() ([]*Product, error)
}


type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int    `json:"quantity"`
	Image    string `json:"image"`
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