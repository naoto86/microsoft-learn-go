package main

import "fmt"

// Employee is sample struct
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

// Person is sample struct
type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

// Employee2 is sample struct
type Employee2 struct {
	Person
	ManagerID int
}

// type Contractor struct {
// 	Person
// 	CompanyID int
// }

func main() {
	employee := Employee{100, "John", "Doe", "Doe's Street"}
	fmt.Println(employee)
	employee.ID = 200
	fmt.Println(employee)
	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)

	employee2 := Employee2{
		Person: Person{
			FirstName: "John",
		},
	}
	employee2.LastName = "Doe"
	fmt.Println(employee.FirstName)
}
