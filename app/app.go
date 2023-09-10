package app

import "net/http"

func Start() {

	// Defined custom mux and register route
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", customers)

	// starting server
	http.ListenAndServe("localhost:8080", mux)
}
