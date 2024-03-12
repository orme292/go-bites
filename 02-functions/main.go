package main

import (
	"fmt"
	"os"
)

// addAndTell accepts two parameters: x and y, both 'int'. The function will print a message and
// the sum of each to the screen.
func addAndTell(x int, y int) {
	fmt.Println("x + y =", x+y)
}

// addNumbers accepts two parameters: x and y, both 'int'. It returns a value of type 'int'.
// A function CANNOT return a value unless it's type is specified in the function signature.
func addNumbers(x int, y int) int {
	return x + y
}

// combineStrings accepts two parameters: a and b, both 'string'. It returns `combined`, a string,
// and `x` an `int`. 'combined' will equal the concatenation of a, a space, and b. `x` will be
// the number 5 and is just used to illustrate a point.
func combineStrings(a string, b string) (combined string, x int) {
	combined = a + " " + b
	x = 5
	return combined, x
}

// returnTheSame accepts a single parameter: x, an 'int'. It returns an 'int' and a 'string'.
// the returned int will be equal to the x parameter and the string will be "we should talk
// about this". No other purpose.
func returnTheSame(x int) (int, string) {
	return x, "we should talk about this"
}

func main() {
	// Print "5 + 5 = 10" to the screen
	addAndTell(5, 5)

	// This calls addNumbers and puts the return value into the 'answer' variable.
	// Then, it will print a string and the answer "The answer is: 3" to the screen.
	answer := addNumbers(1, 2)
	fmt.Println("The answer is: ", answer)

	// This calls combineStrings and puts the return data into the 'result' variable.
	// Then, it will print the result, "Hello World", to the screen.
	result, _ := combineStrings("Hello", "World")
	fmt.Println(result)

	// returnTheSame returns the value of the first parameter and a string.
	// Using the _ allows the code in the returnTheSame to run, while ignoring
	// anything it returns.
	_, _ = returnTheSame(1)

	// Exit the program gracefully.
	os.Exit(0)
}
