package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ayobami6/go_ecom/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	
	t.Run("Should fail if the user is invalid", func (t *testing.T)  {
		payload := types.RegisterUserPayLoad{
			Email:    "invalid",
            Password: "password123",
			FirstName: "first_name",
			LastName:  "last_name",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err!= nil {
            t.Fatal(err)
        }
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest{
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
		
	})
	t.Run("Should Create User Successfully", func (t *testing.T)  {
		payload := types.RegisterUserPayLoad{
			Email:    "ayobamidele001@gmail.com",
            Password: "password123",
			FirstName: "first_name",
			LastName:  "last_name",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err!= nil {
            t.Fatal(err)
        }
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated{
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
		
	})

}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user *types.User) error {
    return nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}
