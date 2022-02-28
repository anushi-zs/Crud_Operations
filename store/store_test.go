package store

import (
	"database/sql"
	goError "errors"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"regexp"
	"testing"
)

func TestEmpbyID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}
	geterr := goError.New("negative Id")
	rows := sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(3, "ANUSHI", "av@gmail.com", "SDE")

	testCases := []struct {
		id            int
		emp           *Employee
		mockQuery     interface{}
		expectedError error
	}{
		{
			id:            3,
			emp:           &Employee{3, "ANUSHI", "av@gmail.com", "SDE"},
			mockQuery:     mock.ExpectQuery("SELECT * FROM employee WHERE Id = ?").WithArgs(3).WillReturnRows(rows),
			expectedError: nil},

		{
			id:            7,
			emp:           nil,
			mockQuery:     mock.ExpectQuery("SELECT * FROM employee WHERE Id = ?").WithArgs(7).WillReturnError(goError.New("no data with given Id")),
			expectedError: sql.ErrNoRows},

		{
			id:            -1,
			emp:           nil,
			mockQuery:     mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM employee WHERE Id = ?")).WithArgs(-1).WillReturnError(geterr),
			expectedError: geterr},
	}

	for _, testCase := range testCases {

		emp, err := EmployeeByID(testCase.id, db)

		if !reflect.DeepEqual(err, testCase.expectedError) {
			t.Errorf("expected error %v got %v", testCase.expectedError, err)
		}
		if !reflect.DeepEqual(emp, testCase.emp) {
			t.Errorf("expected employee %v got %v ", Employee{3, "ANUSHI", "av@gmail.com", "SDE"}, emp)
		}
	}
}

func TestDeletemployee(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}

	deleteErr := goError.New("delete failed")

	cases := []struct {
		desc      string
		id        int
		expecterr error
		mockCall  *sqlmock.ExpectedExec
	}{
		{"Delete success ",
			4,
			nil,
			mock.ExpectExec("delete from employee WHERE Id=?").WithArgs(4).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			"Delete failed",
			7,
			deleteErr,
			mock.ExpectExec("delete from employee WHERE Id=?").WithArgs(7).WillReturnError(deleteErr),
		},
		{"Negative Id ",
			-1,
			goError.New("negative Id"),
			mock.ExpectExec("delete from employee WHERE Id=?").WithArgs(-1).WillReturnError(goError.New("negative Id")),
		},
	}

	for _, tc := range cases {
		err := Deletemployee(tc.id, db)
		if !reflect.DeepEqual(err, tc.expecterr) {
			t.Errorf("expected error %v got %v", tc.expecterr, err)
		}
	}
}
func TestCreateemployee(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}
	createerr := goError.New("create failed")

	tests := []struct {
		desc      string
		id        int64
		expecterr error
		input_emp Employee
		mockCall  *sqlmock.ExpectedExec
	}{
		{
			desc:      "Create Success",
			id:        1,
			expecterr: nil,
			input_emp: Employee{Id: 1, Name: "Anushi", Email: "av@gmail.com", Role: "Engineer"},
			mockCall:  mock.ExpectExec("INSERT INTO employee (Id,Name, Email, Role) VALUES (?,?, ?, ?)").WithArgs(1, "Anushi", "av@gmail.com", "Engineer").WillReturnResult(sqlmock.NewResult(1, 1)),
		},

		{
			desc:      " Create Fail",
			id:        -1,
			expecterr: createerr,
			input_emp: Employee{-1, "Anushi", "av@gmail.com", "Engineer"},
			mockCall:  mock.ExpectExec("INSERT INTO employee (Id,Name, Email, Role) VALUES (?,?, ?, ?)").WithArgs(-1, "Anushi", "av@gmail.com", "Engineer").WillReturnError(createerr),
		},
	}
	for _, tc := range tests {
		emps, err := Createemployee(tc.input_emp, db)

		if emps != tc.id {
			t.Errorf("Expected: %v, Got: %v", tc.id, emps)
		}
		if !reflect.DeepEqual(err, tc.expecterr) {
			t.Errorf("Expected: %v, Got: %v", tc.expecterr, err)
		}
	}
}

func TestEmployeeupdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}
	updaterr := goError.New("update failed")

	tests := []struct {
		desc      string
		expecterr error
		input_emp Employee
		mockCall  *sqlmock.ExpectedExec
	}{
		{
			desc:      "update succes",
			expecterr: nil,
			input_emp: Employee{5, "Aakash", "ak@gmail.com", "Java"},
			mockCall:  mock.ExpectExec("UPDATE employee SET Name = ?, Email=?, Role=? WHERE ID = ?").WithArgs("Aakash", "ak@gmail.com", "Java", 5).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:      "update fail",
			expecterr: updaterr,
			input_emp: Employee{2, "", "", ""},
			mockCall:  mock.ExpectExec("UPDATE employee SET Name=?,Email=?,Role=? WHERE ID = ?").WithArgs("", "", "", 2).WillReturnError(updaterr),
		},
	}
	for _, tc := range tests {
		err := Employeeupdate(tc.input_emp, db)

		if !reflect.DeepEqual(err, tc.expecterr) {
			t.Errorf("Expected: %v, Got: %v", tc.expecterr, err)
		}

	}

}
