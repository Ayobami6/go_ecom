package product

import (
	"net/http"

	"github.com/Ayobami6/go_ecom/types"
	"github.com/Ayobami6/go_ecom/utils"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	store types.ProductStore
}

func NewProductHandler(store types.ProductStore) *ProductHandler {
    return &ProductHandler{store: store}
}

func (h *ProductHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleListProducts).Methods("GET")
    // router.HandleFunc("/products/{id}", h.handleGetProduct).Methods("GET")
    router.HandleFunc("/products", h.handleCreateProduct).Methods("POST")
    // router.HandleFunc("/products/{id}", h.handleUpdateProduct).Methods("PUT")
    // router.HandleFunc("/products/{id}", h.handleDeleteProduct).Methods("DELETE")
}

func (h *ProductHandler) handleListProducts(w http.ResponseWriter, r *http.Request) {
	//... implement product listing logic here
	products, err := h.store.GetProducts()
	if err!= nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }
	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	//... implement product creation logic here
	var payload types.ProductCreatePayLoad
    if err := utils.ParseJSON(r, &payload); err!= nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }
    product := &types.Product{
        Name:        payload.Name,
        Description: payload.Description,
        Price:       payload.Price,
		Quantity: payload.Quantity,
    }
    createdProduct, err := h.store.CreateProduct(product)
    if err!= nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }
    utils.WriteJSON(w, http.StatusCreated, createdProduct)
}