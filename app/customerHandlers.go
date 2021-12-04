package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcelobarreto/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, _ := ch.service.GetAllCustomers(status)
	writeJSONResponse(w, http.StatusOK, customers)
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
