package main

import "fmt"

type Salaried interface {
	getSalary() int
}

type Salary struct {
	basic     int
	insurance int
	allowance int
}

func ( s Salary ) getSalary() int {
	return s.basic + s.insurance + s.allowance
}

type Employee struct {
	firstName, lastName string
	salary Salaried
}

func main() {
	ross := Employee{
		firstName: "Ross",
		lastName:  "Geller",
		salary:    Salary{1100, 50, 50},
	}

	fmt.Println("Ross's  salary is", ross.salary.getSalary())
}
