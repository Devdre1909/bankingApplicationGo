package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name"`
	City string `json:"city"`
}

func main() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{
			Name: "Adegoke Temitope",
			City: "Akure",
		},
		{
			Name: "Toluwanimi Daniel",
			City: "Akure",
		},
		{
			Name: "Ayomide Akindehin",
			City: "Ibadan",
		},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
