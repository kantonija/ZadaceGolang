package main

import "fmt"

func main() {

	Zadatak1()
	Zadatak1pokazivaci()
	Zadatak2()

}

func Zadatak1() {

	// 	Declare two integers, `firstNumber` and `secondNumber`, assign values 20 and 30 to them.
	// Swap values of `firstNumber` and `secondNumber` without using third variable
	// After all, print values of `firstNumber` and `secondNumber`.

	var firstNumber int = 20
	var secondNumber int = 30

	firstNumber, secondNumber = secondNumber, firstNumber

	fmt.Println(firstNumber)
	fmt.Println(secondNumber)
}

func Zadatak1pokazivaci() {

	// preko pokazivača

	var firstNumber *int
	var secondNumber *int

	x := 20
	firstNumber = &x
	y := 30
	secondNumber = &y

	fmt.Println(*firstNumber)
	fmt.Println(*secondNumber)

	firstNumber = &y
	secondNumber = &x

	fmt.Println(*firstNumber)
	fmt.Println(*secondNumber)
}

func Zadatak2() {

	// 	Declare two variables, `firstName` and `lastName` assign them with wanted values.
	// 	Declare constant named `fullname`
	// 	Combine constant and both strings into a full name by concatenating strings with a space in between and print them out.

	var firstName string = "Antonija"
	var lastName string = "Kozul"
	const fullname string = " "
	var ime string = firstName + fullname + lastName

	fmt.Println(ime)
}
