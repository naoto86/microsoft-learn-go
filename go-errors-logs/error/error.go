package main

import (
	"errors"
	"fmt"
)

// Employee is sample struct
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Adress    string
}

func main() {
	employee, err := getInformatin(1001)
	if err != nil {
		return
	} else {
		fmt.Print(employee)
	}
}

// ErrNotFound is sample Error
var ErrNotFound = errors.New("Employee not found!")

func getInformatin(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}
	employee, err := apiCallEmployee(1000)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}
	return employee, nil
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
