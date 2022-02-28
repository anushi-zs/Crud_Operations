package store

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

//Employee struct
type Employee struct {
	Id    int
	Name  string
	Email string
	Role  string
}

var db *sql.DB

//function for getting Id and details of particular employee

func EmployeeByID(id int, db *sql.DB) (*Employee, error) {
	var emp Employee

	if id < 1 {
		return nil, errors.New("negative Id")
	}
	row := db.QueryRow("SELECT * FROM employee WHERE Id = ?", id)
	err := row.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Role)
	if err != nil {
		return nil, sql.ErrNoRows
	}
	return &emp, nil
}

//function for delete employee

func Deletemployee(id int, db *sql.DB) error {

	if id < 1 {
		return errors.New("negative Id")
	}
	result, err := db.Exec("delete from employee WHERE Id=?", id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	//if err != nil {
	//	return errors.New("delete failed")
	//}
	return nil
}

//function for creating new employee

func Createemployee(emp Employee, db *sql.DB) (int64, error) {
	result, err := db.Exec("INSERT INTO employee (Id,Name, Email, Role) VALUES (?,?, ?, ?)", emp.Id, emp.Name, emp.Email, emp.Role)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	//if err != nil {
	//	return -1, errors.New("create failed")
	//}
	return int64(id), nil
}

//function for updating values of  existing employee

func Employeeupdate(emp Employee, db *sql.DB) error {
	_, err := db.Exec("UPDATE employee SET Name = ?, Email=?, Role=? WHERE ID = ?",
		&emp.Name, &emp.Email, &emp.Role, &emp.Id)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}
