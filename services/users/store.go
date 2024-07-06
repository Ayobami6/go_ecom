package user

import (
	"database/sql"
	"fmt"

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
	if err!= nil {
        return nil, err
    }
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err!= nil {
            return nil, err
        }

	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
    err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)
    if err!= nil {
        return nil, err
    }
    return u, nil
}

func (s *Store) CreateUser(u *types.User) error {
	return nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}