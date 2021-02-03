package main

import "fmt"

// CarDetails are ...
type CarDetails struct {
	carName, carModel string
	carPrice          int
	discountAllowed   bool
}

func main() {
	// The var keyword initializes a variable rect. Using dot notation, values are assigned to the struct fields.
	var jaguar CarDetails
	jaguar.carName = "Jaguar"
	jaguar.carModel = "RSport"
	jaguar.carPrice = 34000
	jaguar.discountAllowed = true
	fmt.Println(jaguar)

	// Creates an instance of CarDetails struct by using a struct literal and assigning values to the fields of the struct.
	bmw := CarDetails{"BMW", "3 Series", 33000, false}
	fmt.Println(bmw)
}
