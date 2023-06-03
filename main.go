package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
