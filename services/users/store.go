package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Ayobami6/go_ecom/types"
)

// repository pattern

type Store struct {
	db *sql.DB

}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query(`SELECT * FROM users WHERE email = ?`, email)
	fmt.Println(rows)
	if err!= nil {
		log.Fatal("This is the error", err)
        return nil, err
    }
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows, u)
		if err!= nil {
            return nil, err
        }

	}
	// user not found
	if u.ID == 0 {
		return nil, errors.New("user not found")
	}
	// I found a user
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows, u *types.User) (*types.User, error) {
    err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)
    if err!= nil {
        return nil, err
    }
    return u, nil
}

func (s *Store) CreateUser(u *types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query(`SELECT * FROM users WHERE id = ?`, id)
	if err!= nil {
        return nil, err
    }
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows, u)
		if err!= nil {
            return nil, err
        }

	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}