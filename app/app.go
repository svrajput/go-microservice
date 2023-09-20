package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/svrajput/go-microservice/domain"
	"github.com/svrajput/go-microservice/service"
)

func Start() {

	SanityCheck()

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

	serverAdd := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	srvString := fmt.Sprintf("%s:%s", serverAdd, serverPort)

	// starting http server
	log.Panic(http.ListenAndServe(srvString, router))

}

func SanityCheck() {
	//TODO - get all environment variables
	envVar := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, k := range envVar {
		if os.Getenv(k) == "" {
			log.Fatalf(fmt.Sprintf("Environment Variable is missing %s", k))
		}
	}
}
