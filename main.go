package main

import (
	"bank/repository"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:keep1234@tcp(127.0.0.1)/banking")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)

	_ = customerRepository

	// customers, err := customerRepository.GetAll()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)
	_ = customerRepository

	customer, err := customerRepository.GetById(2000)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
