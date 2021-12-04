package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcelobarreto/banking/errs"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func scanRow(row *sql.Row, c *Customer) *errs.AppError {
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("customer not found")
		} else {
			return errs.NewUnexpectedError("unexpected database error")
		}
	}
	return nil
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	rows, err := d.client.Query(findAllSql)

	if err != nil {
		return nil, errs.NewUnexpectedError("Error while querying customers table " + err.Error())
	}
	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

		if err != nil {
			return nil, errs.NewUnexpectedError("Error while scanning customer table " + err.Error())
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	var c Customer
	row := d.client.QueryRow("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?", id)
	err := scanRow(row, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "marcelo:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
