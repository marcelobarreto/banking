package domain

import (
	"database/sql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcelobarreto/banking/errs"
	"github.com/marcelobarreto/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func scanRow(row *sql.Rows, c *Customer) *errs.AppError {
	if row == nil {
		return &errs.AppError{Code: http.StatusNotFound, Message: "customer not found"}
	}

	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customers table " + err.Error())
			return errs.NewUnexpectedError("unexpected database error")
		}
	}
	return nil
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?"
		rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := scanRow(rows, &c)

		if err != nil {
			logger.Error(err.Message)
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	var c Customer
	rows, _ := d.client.Query("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?", id)
	if rows.Next() {
		err := scanRow(rows, &c)
		if err != nil {
			logger.Error(err.Message)
			return nil, err
		}

		return &c, nil
	}

	return nil, &errs.AppError{Code: http.StatusNotFound, Message: "customer not found"}
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "marcelo:password@tcp(localhost:3306)/banking")
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
