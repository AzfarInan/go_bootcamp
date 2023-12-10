package main

import "fmt"

func main() {
	age := 29
	fmt.Println("My age is", age)

	s := "Learning Go!"
	fmt.Println(s)

	_ = "Learning Go!"
	name := "Nigel"
	name = "Commons"
	fmt.Println(name)

	/// Multiple Declarations
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	car, year := "BMW", 2018
	_ = year

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	j, i = i, j // Swaping variables
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}
