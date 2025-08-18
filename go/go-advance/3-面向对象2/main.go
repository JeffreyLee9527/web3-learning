package main

import "fmt"

func main() {
	e := Employee{121, Person{"AAA", 18}}
	e.PrintInfo()
}

type Person struct {
	Name string
	Age  int8
}

type Employee struct {
	EmployeeID int8
	Person
}

func (e Employee) PrintInfo() {
	fmt.Printf("Employee  : %+v", e)
}
