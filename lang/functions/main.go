package main

import (
  "fmt"
)

// Entry point of every program:
// - main function in package main that takes no parameters and returns no value
func main() {

  // Functions
  // Starts with func keyword
  // PascalCase or camelCase to determine visibility
  // Compiler enforces first curly brace at end of func line and last curly brace on its own line

  sayMessage("Hello Go!", 22) // Hello Go! 22
  sayGreeting("Hello", "Bob") // Hello Bob

  // Pass by value
  greeting := "Hello"
  name := "Bob"
  // Hello Bob
  // Ted
  passedByValue(greeting, name)
  fmt.Println(name) // Bob

  // Passing pointers
  // Can be much more efficient with large data sets
  // Note that slices and maps are internal pointers so they will always by passed by reference

  // Hello Bob
  // Ted
  passedByPointer(&greeting, &name)
  fmt.Println(name) // Ted

  // Variadic parameter
  msg := "The sum is "
  sumInPlc(msg, 1, 2, 3, 4, 5) // The sum is 15

  // Return value
  theSum := sum(msg, 1, 2, 4)
  fmt.Println(msg, theSum) // The sum is 7 

  // Return local variable as pointer
  pSum := sumPointer(11, 11)
  fmt.Println("The sum is ", *pSum) // The sum is 22

  // Return named values
  // Not commonly used
  nSum := sumNamed(5, 4)
  fmt.Println("The sum is ", nSum) // The sum is 9

  // Multiple return values
  d, err := divide(5.0, 0.0)
  if err != nil {
    fmt.Println(err) // Cannot divide by zero
    // return // will exit entire program
  }
  fmt.Println(d) // 0

  // Anonymous functions
  // Hello Go! 
  anonymousFn()

  // Functions as types or as variables
  f := func() {
    fmt.Println("anonymous assinged to f")
  }
  f() // anonymous assinged to f

  var fn func() = func() {
    fmt.Println("anonymous assinged to fn")
  }
  fn() // anonymous assinged to fn

  var fnVerbose func(float64, float64) (float64, error)
  fnVerbose = func(a, b float64) (float64, error) {
    return a + b, nil
  }
  r, err := fnVerbose(12.33, 45.6)
  fmt.Println(r, err) // 57.93 <nil>

  // Methods
  4:48
}

func sayMessage(msg string, idx int) {
  fmt.Println(msg, idx)
}

// Syntax sugar when same types
func sayGreeting(greeting, msg string) {
  fmt.Println(greeting, msg)
}

func passedByValue(greeting, name string) {
  fmt.Println(greeting, name)
  name = "Ted"
  fmt.Println(name) // Ted
}

func passedByPointer(greeting, name *string) {
  fmt.Println(*greeting, *name)
  *name = "Ted"
  fmt.Println(*name) // Ted
}

// Variadic paramter
// Can only be one
// Must be last parameters
func sumInPlc(msg string, values ...int) {
  fmt.Println(values)
  result := 0
  for _, v := range values {
    result += v
  }
  fmt.Println(msg, result)
}

// Return value
func sum(msg string, values ...int) int {
  fmt.Println(values)
  result := 0
  for _, v := range values {
    result += v
  }
  return result
}

// Return a pointer
// In most languages it's not safe to return reference declared on stack
// because it will be destroyed when function terminates
// Once function terminates all memory stored on execution stack is freed.
// However, Go recognizes you're returning a pointer and will automatically promte it to the heap memory
func sumPointer(values ...int) *int {
  fmt.Println(values)
  result := 0 // result is declared on execution stack of this function
  for _, v := range values {
    result += v
  }
  return &result
}

// Named return values
// Syntax sugar
func sumNamed(values ...int) (result int) {
  fmt.Println(values)
  for _, v := range values {
    result += v
  }
  return // no need to specify as it is implicit
}

func divide(a, b float64) (float64, error) {
  if a == 0.0 || b == 0.0 {
    return 0.0, fmt.Errorf("Cannot divide by zero")
  }
  return a /b, nil
}

func anonymousFn() {
  func() {
    fmt.Println("Hello Go!")
  }() // Note the braces to immediately envoke the function

  // Pass a value to the inner scope
  // Especially due to concurrency issues
  func(msg string) {
    fmt.Println(msg)
  }("Goodbye Go!")
}

