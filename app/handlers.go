package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcelobarreto/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		writeXMLResponse(w, http.StatusOK, customers)
	} else {
		writeJSONResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeJSONResponse(w, err.Code, err.AsMessage())
	} else {
		writeJSONResponse(w, http.StatusOK, customer)
	}
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post received")
}

func writeJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func writeXMLResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(code)
	if err := xml.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
