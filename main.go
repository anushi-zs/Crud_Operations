package main

import (
	"database/sql"
	"fmt"
	"github.com/anushi/newbatch/Crud-operations/store"
	"log"
)

func main() {

	//Steps for sql connection
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected")
	}
	//defer db.Close()

	//Call for getting employee id
	emp, err := store.EmployeeByID(3, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Employee found: %v\n", emp)

	//Delete Employee by id
	err = store.Deletemployee(4, db)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Deleted Succesfully")
	}

	//Call for create employee
	empID, err := store.Createemployee(store.Employee{1, "Anushi", "av@gmail.com", "Engineer"}, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added employee: %v\n", empID)

	//Call for update
	err = store.Employeeupdate(store.Employee{5, "Aakash", "ak@gmail.com", "Java"}, db)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Updated Succesfully")
	}
}
