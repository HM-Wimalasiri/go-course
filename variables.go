package main

import "fmt"
var middleName="Cane"
func main(){
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// count :=10
	// lastName :="Smith"

	// middleName:="Alex"

	fmt.Println(middleName)

	// Default Values
	// Numeric Types : 0
	// Boolean Types : False
	// String Types : ""
	// Pointers, slices, maps, functions and struc: nil

	// --- Scope
	printName()
}

func printName(){
	firstName :="Micheal"
	fmt.Println(firstName)
}