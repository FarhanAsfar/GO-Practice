package main

import (
	"fmt"
)

//A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type

type Staff struct {
	name     string
	salary   int
	currency string
}

// here, displaySalary() mehtod has Staff as thereceiver type
func (e Staff) displaySalary() {
	fmt.Printf("salary of %s is %s%d ", e.name, e.currency, e.salary)
}

func Methods() {
	emp1 := Staff{
		name:     "Sam Adolf",
		salary:   40000,
		currency: "$",
	}

	emp1.displaySalary() //calling displaySalary() method of Employee type
}
