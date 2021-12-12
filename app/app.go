package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/marcelobarreto/banking/domain"
	"github.com/marcelobarreto/banking/logger"
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

	dbClient := getDBClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDB(dbClient)

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func getDBClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
