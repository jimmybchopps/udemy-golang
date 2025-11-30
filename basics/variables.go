package main

import "fmt"

// middleName := "Cane"
var middleName = "Cane"

func main() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"
	
	// Type inference
	// count := 10
	// lastName := "Smith"

	// Default values
	// Numeric Types: 0
	// Boolean Types: false
	// String Types: ""
	// Pointers, Slices, maps, functions, and structs: nil

	// ------ SCOPE
	// := Gopher notation can only be used in local scopes within funcitons
	middleName = "Mayor"
	fmt.Print(middleName)

}

func printName() {
	firstName := "Michael"
	fmt.Println(firstName)
}