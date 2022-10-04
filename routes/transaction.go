package routes

import (
	"_waysbook/handlers"
	"_waysbook/pkg/middleware"
	"_waysbook/pkg/mysql"
	"_waysbook/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
  transactionRepository := repositories.RepositoryTransaction(mysql.DB)
  h := handlers.HandlerTransaction(transactionRepository)

  r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
  r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
  r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
  // r.HandleFunc("/transaction/{id}", h.UpdateTransactions).Methods("PATCH")
  r.HandleFunc("/transaction", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
  r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE")
	r.HandleFunc("/transaction-status", middleware.Auth(h.FindbyIDTransaction)).Methods("GET")
  r.HandleFunc("/notification", h.Notification).Methods("POST")
}
