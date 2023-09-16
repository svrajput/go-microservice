package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/svrajput/go-microservice/domain"
	"github.com/svrajput/go-microservice/service"
)

func Start() {

	// Defined custom router and register route
	//router := http.NewServeMux()
	router := mux.NewRouter()

	//wire
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// Customer Routes
	router.HandleFunc("/customers", ch.customers).Methods(http.MethodGet)

	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/createCustomer/", createCustomer).Methods(http.MethodPost)
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)

	// starting http server
	http.ListenAndServe("localhost:8080", router)

}
