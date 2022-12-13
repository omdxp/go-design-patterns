package main

import "fmt"

type Address struct {
	Suite                        int
	StreetAddress, City, Country string
}

type Employee struct {
	Name   string
	Office Address
}

var mainOffice = Employee{
	"", Address{0, "123 East Dr", "London", "UK"},
}

var auxOffice = Employee{
	"", Address{0, "66 West Dr", "London", "UK"},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := *proto
	result.Name = name
	result.Office.Suite = suite
	return &result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}

// Prototype Factory is a method that creates a copy of the object by using a factory.
