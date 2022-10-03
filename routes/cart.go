package routes

import (
	"_waysbook/handlers"
	"_waysbook/pkg/middleware"
	"_waysbook/pkg/mysql"
	"_waysbook/repositories"

	"github.com/gorilla/mux"
)

func CartRoute(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", h.FindCart).Methods("GET")
	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart-id", middleware.Auth(h.FindCartsByTransaction)).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.UpdateCart)).Methods("PATCH")
}