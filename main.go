package main

import (
	"_waysbook/database"
	// "_waysbook/handlers"
	"_waysbook/pkg/mysql"
	"_waysbook/routes"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	errEnv := godotenv.Load()
    if errEnv != nil {
      panic("Failed to load env file")
    }
	
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	// fmt.Println("server running localhost:5000")
	// http.ListenAndServe("localhost:5000", r)

	var port = "5000"
	fmt.Println("server running localhost:"+port)
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}