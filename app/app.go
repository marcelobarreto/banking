package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcelobarreto/banking/domain"
	"github.com/marcelobarreto/banking/service"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":3000", router))
}
