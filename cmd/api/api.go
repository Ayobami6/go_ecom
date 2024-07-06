package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Ayobami6/go_ecom/order"
	"github.com/Ayobami6/go_ecom/services/cart"
	"github.com/Ayobami6/go_ecom/services/product"
	root "github.com/Ayobami6/go_ecom/services/root"
	user "github.com/Ayobami6/go_ecom/services/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1/").Subrouter()

	userStore := user.NewStore(s.db)
	productStore := product.NewProductStore(s.db)

	userHandler := user.NewHandler(userStore)
	rootHandler := root.NewRootHandler()
	productHandler := product.NewProductHandler(productStore)
	userHandler.RegisterRoutes(subrouter)
	rootHandler.RegisterRoutes(subrouter)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartStore := cart.NewCartStore(s.db)

	cartHandler := cart.NewCartHandler(cartStore, productStore, userStore,  orderStore)
	cartHandler.RegisterRoutes(subrouter)
	log.Println("Listening on: ", s.addr)

	return http.ListenAndServe(s.addr, router)


}