package routes

import (
	"_waysbook/handlers"
	"_waysbook/pkg/middleware"
	"_waysbook/pkg/mysql"
	"_waysbook/repositories"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
  bookRepository := repositories.RepositoryBook(mysql.DB)
  h := handlers.HandlerBook(bookRepository)

  r.HandleFunc("/books", h.FindBooks).Methods("GET")
  r.HandleFunc("/book/{id}", h.GetBook).Methods("GET")
  r.HandleFunc("/book", middleware.Auth(middleware.UploadPDF(middleware.UploadFile(h.CreateBook)))).Methods("POST")
//   r.HandleFunc("/book/{id}", h.UpdateBook).Methods("PATCH")
  r.HandleFunc("/book/{id}", h.DeleteBook).Methods("DELETE")
}