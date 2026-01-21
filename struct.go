package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	address Address
}

func Struct() {
	emp1 := Employee{
		firstName: "Farhan",
		lastName:  "Asfar",
		age:       25,
		salary:    000,
	}

	//creating struct without specifying field names, dont use this syntax
	emp2 := Employee{"Thomas", "Paul", 29, 3000}

	fmt.Println("First Name: ", emp1.firstName)
	fmt.Println("Last Name: ", emp1.lastName)
	fmt.Println("Age: ", emp1.age)
	fmt.Println("Salary: ", emp1.salary)
	fmt.Println(emp2)

	//nested struct
	p := Person{
		name: "Alice",
		age:  40,
		address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("City: ", p.address.city)
	fmt.Println("Address: ", p.address.state)

}
