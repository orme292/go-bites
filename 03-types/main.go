package main

import (
	"fmt"
	"reflect"
)

func main() {
	// explicit declaration with a type
	var myNumbers int8

	// short variable declaration with :=
	yourLetters := "letters for you"

	// print types to the screen using the 'reflect' package
	fmt.Println(reflect.TypeOf(myNumbers))   // prints: int8
	fmt.Println(reflect.TypeOf(yourLetters)) // prints: string

	// assign values to multiple variables using short variable declaration
	myNum1, myNum2, myFloat1 := 50, 75, 200.1

	// Using fmt.Printf, print the names and values of each variable to the screen
	fmt.Printf("\nVariable Set 1:\n\tmyNum1: %d\n\tmyNum2: %d\n\tmyFloat1: %f\n", myNum1, myNum2, myFloat1)

	// declare variables in a single var block, with parenthesis (), and then
	// assign values separately using =
	// finally, print the names and values to the screen using fmt.Printf
	var (
		age    int
		name   string
		height float32
	)
	age = 29
	name = "Marcus"
	height = 70.999
	fmt.Printf("\nVariable Set 2:\n\tname: %s\n\tage: %d\n\theight (inches): %.3f\n", name, age, height)

	myString := "string!"
	myString := "different string!"
}
