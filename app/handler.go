package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip" xml:"zip"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world rest response")
}

func customers(w http.ResponseWriter, r *http.Request) {
	customers := []customer{
		{"Jason", "Dallas", "1111111"},
		{"John", "Dallas", "222222"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func getCustomers(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	cid := params["customer_id"]

	fmt.Fprintln(w, cid)

}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "post request received for new customer.")
}
