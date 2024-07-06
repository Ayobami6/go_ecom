package cart

import (
	"net/http"

	"github.com/Ayobami6/go_ecom/services/auth"
	"github.com/Ayobami6/go_ecom/types"
	"github.com/Ayobami6/go_ecom/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type CartHandler struct {
	store types.CartStore
	pStore types.ProductStore
	orderStore types.OrderStore
	userStore types.UserStore
}

func NewCartHandler(store types.CartStore, ps types.ProductStore, us types.UserStore, os types.OrderStore) *CartHandler {
    return &CartHandler{store: store, pStore: ps, orderStore: os, userStore: us}
}


func (c *CartHandler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/cart/checkout", auth.Auth(c.handleCheckout, c.userStore)).Methods("POST")
}

func (c *CartHandler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := 0
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errs := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errs)
		return
	}

	productIds, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get products
	products, err := c.pStore.GetProductsByID(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := c.createOrder(products, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})



	
}

