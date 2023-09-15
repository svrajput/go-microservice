package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// Defined custom router and register route
	//router := http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)

	// Customer Routes
	router.HandleFunc("/customers", customers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/createCustomer/", createCustomer).Methods(http.MethodPost)

	// starting http server
	http.ListenAndServe("localhost:8080", router)

}
