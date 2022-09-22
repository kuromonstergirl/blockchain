package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person
	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	//printinG
	printPerson(pers1)
	printPerson(pers2)
}
func printPerson(p Person) {
	// code to be executed
	fmt.Printf("Name of person   : %s/n", p.name)
	fmt.Printf("age of person    : %d\n", p.age)
	fmt.Printf("job of person    : %s/n", p.job)
	fmt.Printf("Salary of person : %d\n", p.salary)

}
