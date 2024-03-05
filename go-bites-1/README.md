# Bite-sized Go: Part 1 - The Basics

So, you already know how to code. Maybe it's not what you do all day - or maybe it is what you do all day. But, you get it: variables, functions, loops, etc etc. Plus… you're busy, but you still want to learn Go. Then, maybe, this is for you.

You've heard about Go… everyone has heard about Go. You know that it's crowning achievement is that it was both created by Google and not (yet) cancelled by Google. Which is probably good, since it powers some of the most important containerization software out there (Docker, Kubernetes).

This series won't go so deep that you'll be able to write you're own containerization software, but there's a strong possibility that you'll confidently write you're first Hello World program in 4–6 weeks (kidding).

Honestly, my goal is to break down the basic elements of Golang. I'll touch on the important stuff that, hopefully, will help you get started. Just remember that when I say bite-sized, I mean bite-sized. These are short reads. With quick visuals of code blocks that, I'm hoping, make it possible to read quickly and still come away with something.

---

## Why Go

Go's founding goal was to simplify software development at Google. Unlike the languages it sought replace (C++, Java, Python), Go was meant to be fitter, happier, and more productive. Official word is that the goal was reasonably met. Go is simple and incredibly powerful. It's an excellent language that combines Python-like readability with C++-like capability.

By far, Go's best feature is the `go fmt` command: It takes whatever junk you just wrote and formats it the _Go way_. Whether you use VSCode, Zed, or GoLand, this should happen automatically. If you use Vim, figure it out: that's why you use Vim.

To move forward, you have to install Go. It's easy! Download it at [go.dev](https://go.dev/dl/).

## 1: Your First Go File
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Pretty simple, right? We've all the basics here: `package main` defines this file to be part of the `main` package. If you're writing a program that will be executed independently (i.e. by compiling a binary and running it on the command line), then this is where execution will start. If you're writing a library (something someone else can import into their program), I can't think of a reason why you would need a package main. More on that later.

```go
import "fmt"
```

Imports are handled after the package declaration. Every go source file declares the imports needed for that source file. `import "fmt"` tells Go to import the `fmt` package - part of the standard library. The `fmt` package mainly handles inputs and outputs, like printing a line on the screen. Go doesn't like it when you import packages and don't use them. The Go formatter `go fmt` will, with a single exception, always remove unused imports.

```go
func main() {
    fmt.Println("Hello, World!")
}
```

Our main function is where code execution will begin, but only if its part of `package main`. If you write a `func main()` in a package named `math`, then it won't be executed automatically. You'd have to call it some other way.

```go
fmt.Println("Hello, World!")
```

For our Hello World print out, we reference the package `fmt` and use the `Println` function in that package. `Println` only requirement is the string to print on the screen: Hello, World!

Also note the curly braces `{}`. Functions, if-statements, and for- loops are grouped using the curlies. I really appreciate this. My pinky's fingernail doesn't, but I do.

```go
func main() {
  if 1 == 1 {
    if 2 == 2 {
      if 3 == 3 {
        if 4 == 4 {
          if 5 == 5 {
            fmt.Println("It's true!")
          }
        }
      }
    }
  }
}
```

For the Python folks, you'll love the clarity that the `{}` braces provide. However, you'll hate it when your function ends with six closing braces.

That's it - I really meant it when I said bite-sized (and this was a little longer than I intended). Next time, we'll take a short look at functions.

## Running Go Code

Start by creating a new directory named "go-bites-1" (`mkdir go-bites-1`) and then creating a new `main.go` file, either with your code editor or with the `touch main.go` command.

Copy and paste the top most code block in this article into your new `main.go` file and then run it from the command line by typing `go run main.go`. You'll see the following output:

```bash
$ go run main.go

Hello, World!
```

## Extra Credit

You can override the name of an imported package if it interferes with your own functions or module names (or some other conflict comes up). Just add the new name before the import.

```go
package main

import theformatpackage "fmt"

func main() {
    theformatpackage.Println("Hello, World!")
}
```

When importing multiple packages, remember to surround them with parenthesis. The `import` keyword can only be used once!

```go
package main

import (
    "os"
    theformatpackage "fmt"
)

func main() {
    theformatpackage.Println("Hello, World!")
    os.Exit(0)
}
```

See the open-source s3 tool, s3packer, derived from the projects I mention in my posts: [https://www.github.com/orme292/s3packer](https://www.github.com/orme292/s3packer)

Follow the series with the repo: [go-bites @ github](https://github.com/orme292/go-bites)