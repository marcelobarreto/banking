package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/marcelobarreto/banking/domain"
	"github.com/marcelobarreto/banking/service"
)

func sanityCheck() {
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal("DATABASE_URL env variable was not set")
	} else if os.Getenv("PORT") == "" {
		log.Fatal("PORT env variable was not set")
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
