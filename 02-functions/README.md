# Bite-sized Go: Part 2 -- Functions
You already know them and you already use them, so this should be cake. They're the smallest and most important unit in Go, which is why I chose to talk about them before types. But, you already know about types somewhat, right?
Previously, I touched on the basic layout of a Hello World program in Go - the main package, importing libraries, and the main() function. So, let's talk more about those functions. We need to address the passing of and returning of parameters.

## Our program
```go
package main

import "os"

func main() {
  os.Exit(0)
}
```
Nothing to see here, right? You know `package main` and `import "os"`. The `main()` function isn't new, either. `os.Exit(0)` might be new to you, but this just calls Exit from the os package, which causes the program to exit normally. You can use `os.Exit(1)` to exit and return an error to the shell.

## #2 Your First Go function
```go
package main

import "os"

func addAndTell(x int, y int) {
  fmt.Println("x + y =", x+y)
}

func main() {
  addAndTell(5, 5)
  os.Exit(0)
}
```

Above, I added a new function named `addAndTell`. This function asks for two parameters: `x` and `y`, both `int` (integers). The `fmt.Println` command will print a string, including the sum of both parameters, to the screen. Lastly, Inside of our `main()` function, we call the new function `addAndTell(5, 5)` before exiting. When we run the code, we'd get the following output:

```text
$ go run main.go

x + y = 10
```

Note that each input parameter (x int, y int) has a data type -- this is required. Go is a strongly typed language; every variable must be assigned a data type.

## Returning Data
```go
func addNumbers(x int, y int) int {
  return x+y
}
```

To return data, a type must be specified after the input parameters. In the case above, addNumbers will return data of type int. Without specifying a data type, your code will not run when you try to return data.

## Naming Return Parameters

All input parameters must be named, but return parameters aren't required to have names. Though, you must specify types for your return data.

```go
func combineStrings(a string, b string) (combined string, x int)) {
  combined = a + " " + b
  x = 5
 return combined, x
}
```

In the above code, I have two return parameters: `combined` and `x`, each a different type. More often than not, I have found that it's easier to use named return parameters… especially when dealing with more than one.

```go
func returnTheSame(x int) (int, string) {
  return x, "we should talk about this"
}
```

The above example will return an int and a string value. Since they aren't named, you must remember to return them in the same type order specified in the function signature.

## Ignoring Return Data

Some functions return a ton of information, maybe more than you need. If that's the case, you can chose to ignore or omit some of it.

```go
func returnTooMuch() (x int, y int, s string) {
  return 1, 2, "pointless"
}

func main() {
  a, b, _ := returnMany()
}
```
  
Use `_` in place of any value you want to ignore. If you just want the code in a function to run, you can ignore every single return value. Go calls the `_` The Blank Identifier.

## Unexported and Exported Functions

```go
func RunMe() {
  fmt.Println("I'm free!")
}

func cantRunMe() {
  fmt.Println("I'm a prisoner.")
}
```

This will become more important as you continue to develop with Go. When creating a package meant to be imported into other people's code, you can make functions either available or unavailable to them. Functions that start with a capital letter func RunMe() are available outside of your package. Functions that start with a lowercase letter func cantRunMe() are NOT available outside of your package. This also applies to variables and type definitions. For now, no need to worry too much about this. Still, I'd say this is one of the stranger things that Go does, though it is simple.

## Strict Typing

To illustrate how important it is to manage your data types, let's tell our function to return the string "apples" instead of the a number.

```go
func addNumbers(x int, y int) int {
 return "apples"
}
```

Our program will not run or compile.

```text
$ go run main.go

./main.go: cannot use "apples" (untyped string constant) as int value in argument to addNumbers
```

And that's a wrap for the section on functions. It's a lot to take in, but it's mainly just rules and requirements. Next time, we'll attack those types.

---

## Extra Credit

When two or more consecutive parameters are of the same type, you only need to specify the type once, rather than for each. This can happen multiple times in a function's signature.

```go
func addNumbers(x, y int) int {
    return x+y
}
```

## In Summary

- Every input parameter must have a name and a type: `func addNumbers(x int, y int)`
- Return parameters can be named `func addNumbers(x int, y int) (z int)` or unnamed `func addNumbers(x int, y int) int`.
- When return data is not named, remember to return the data in the same type order as it was specified in the signature.
- Ignore return values by using `_` when calling the function `num, _, string = getData()`.
- Every single variable in Go must be typed. Using the wrong data in a typed variable will cause prevent your code from running and compiling.
- Exported functions begin with a capital letter `func ForEveryone()`
- Unexported functions being with a lowercase letter `func justForYou()`

---

Follow along, or see the series, in the go-bites repo.

If you enjoyed this article, please consider following me or showing some support with a clap.