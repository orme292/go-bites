# Bite-sized Go: Part 3 — Basic Types
> Types are one of the most important elements of any language. Maybe you’ve written code in Python and you’ve heard people talk about types. Maybe you know them personally. Either way, when it comes to Go, there’s no ignoring them.

Go is a statically typed language, so the types of all variables must be known at compile time — read: every variable must be assigned a type. Go has the usual type suspects: int, bool, and string. Then. there’s the lesser known types, like rune and interface. I’ll talk about the tough ones later. For now, let’s talk about the basic types and the different ways to assign them.

## 1. Numeric Types
These are the number types — they’ll store integers, floating point numbers, and some complex types. The primary difference between them is the size, whether the numbers have a decimal, and whether they can hold signed or unsigned numbers. Here’s a few common ones with a little description (read more).

- `int`: the size of an int depends on the architecture of your system, but in most cases it’ll be 64-bit. An int will hold nearly any whole number you need it to, both “negative” and “positive”
- `uint`: the unsigned integer has a minimum value of 0, and a maximum value of 1844674407370955615, so no negative values.
- `int8`: 8-bits of numerical power — holds any value between -128 and 127
- `uint8`: 8-unsigned-bits of numerical power — holds any value between 0 and 255.
- `float32`/`float64`: these types will store numbers with a decimal point… with varying degrees of precision.

## 2. String Type
`String` types will store, you guessed it, letters and words (two words to be exact, but that’s another story). In reality, the stored data is a sequence of bytes, but all you need to care about for now is: a variable declared as a string will hold characters.

## 3. Boolean Type
A `boolean` type stores a value of `true` or `false`. That’s it.

## 4. The Others
While you can get pretty far with the basic types above, you’ll eventually need something with a bit more power. That’s where the `slice`, `struct`, `pointer`, and `interface` types come in to play. They’ll each be their own article, I think… I don’t want this to drag on.

## Variable Declaration & Assignment
In Go, you primarily have two methods at your disposal for declaring variables.

```go
package main

import (
"fmt"
"reflect"
)

func main() {
// explicit declaration
var myNumbers int8

// short variable declaration with :=
yourLetters := "letters for you"

fmt.Println(reflect.TypeOf(myNumbers))   // prints: int8
fmt.Println(reflect.TypeOf(yourLetters)) // prints: string
}
```

See this in action: https://goplay.tools/snippet/l3_kzBKfAQY

The `var` keyword: You can explicitly declare a variable and assign it a type, as shown in the first example above: `var myNumbers int`. In this line, the code declares a variable named `myNumbers` and specifies that it will hold data only of type `int8`.

The `:=` operator: Go also allows inferential assignment with the short variable declaration operator `:=`. Get comfortable with it now… because you’ll be using this operator liberally. Though, you will lose some specificity. If you expect `myNumbers := 45` to produce a variable typed as an `uint8`, you’ll be disappointed. In this case, Go will assign `myNumbers` an `int` type by default.

It’s also important to remember that the short variable declaration operator `:=` can only be used for new variables.

```go
myString := "string!"
myString := "different string!"
```
The above code will result in a compile error. Since myString was created with := in the first line, Go will attempt to create it again in the second line and fail because it already exists, resulting in the following error:

```text
./main.go: no new variables on left side of :=
```

You cannot assign a new value to an existing variable using `:=`, it can only be used when creating a new variable. The correct way to assign a new value to an existing variable is with the `=` operator.

```go
myString := "string!"
myString = "different string!"
```

## Extra Credit
```go
myNum1, myNum2, myFloat1 := 50, 75, 200.1
name, color, age := "andrew", "green", 100
```
^ You can use short variable declarations := to assign multiple variables of the same or different type all at once.

```go
var num1, num2, num3 int
```
^ Go also allows you to explicitly declare multiple variables of the same type in a single line.

```go
var(
age int
name string
debt float32
)
```
^ Similar to the import keyword, you can declare multiple variables by wrapping them in parenthesis var().

## In Summary

- Variables can be created with the `var` keyword, followed by the variable name and type: `var firstName string`
- Variables can be created and assigned a value with the short variable assignment operator `:=`: `firstName := "Marcus"`.
- Variable types will be inferred when using `:=`. Use the explicit method, with `var`, in order to assign a specific numerical type to your variable.
- You **cannot** use `:=` to assign a new value to an existing variable, doing so will result in a compile error.
- You can assign a new value to an existing variable with `=` equals: `firstName = "Mike"`.
- Signed types, like `int`, `int8`, and `int16`, can hold both negative and positive whole numbers.
- Unsigned types, like `uint`, `uint8`, and `uint16`, can only hold zero and positive whole numbers.
- Floating point types, like `float32`, can hold numbers with a decimal point.

---

Follow along, or see the series, in the go-bites repo.

If you enjoyed this article, please consider following me or showing some support with a clap.