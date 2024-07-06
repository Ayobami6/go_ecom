package user

import (
	"errors"
	"net/http"

	"github.com/Ayobami6/go_ecom/services/auth"
	types "github.com/Ayobami6/go_ecom/types"
	"github.com/Ayobami6/go_ecom/utils"
	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

// type handle for dependencies injection

type Handler struct {
	store types.UserStore
}

// func to create a new Handler

func NewHandler(store types.UserStore) *Handler {
    return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}


func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
    //... implement login logic here
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get json Payload
	var payload types.RegisterUserPayLoad
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// validate
	if err := utils.Validate.Struct(payload); err != nil {
		errs := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errs)
		return
	}
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
        utils.WriteError(w, http.StatusConflict, errors.New("user with this email already exists"))
        return
    }
	hashedPassword, hasErr := auth.HashPassword(payload.Password)
	if hasErr!= nil {
        utils.WriteError(w, http.StatusInternalServerError, hasErr)
        return
    }
	err = h.store.CreateUser(&types.User{
		Email:    payload.Email,
        Password: hashedPassword,
		FirstName: payload.FirstName,
		LastName: payload.LastName,
	})
	if err!= nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
    }
	utils.WriteJSON(w, http.StatusCreated, "User registered successfully")

}