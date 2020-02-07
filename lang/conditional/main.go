package main

import (
  "fmt"
)

func main() {

  // if
  fmt.Println("\nif statements:")

  if true {
    fmt.Println("Yes sir")
  }

  // Initializer/Comma ok syntax
  aMap := map[string]int{
    "aKey": 22,
  }
  // Note: val will only be available in the if block
  if val, ok := aMap["aKey"]; ok {
    fmt.Println("the key exists, the value is: ", val)
  }
  // ERROR if try to use val outside of if block
  // ./main.go:24:37: undefined: val 
  // fmt.Println("val out of scope: ", val)

  // Operators
  // > >= < <= == !=
  // Not operator: !
  // Short circuit operators: && ||
  a := 2
  if (a > 0) {
    fmt.Println("a is greater than 0")
  }

  if (!returnTrue()) {
    fmt.Println("it didn't return true")
  } else {
    fmt.Println("it returned true")
  }

  if a == 999 {
    fmt.Println("This won't print")
  } else if a != 2 {
    fmt.Println("This won't print")
  } else {
    fmt.Println("This prints")
  }

  // switch
  fmt.Println("\nswitch statements:")
  // Go doesn't fall throw, but can use multiple test on each case
  // break is implicit so it's not needed unlike other c languages where fallthrow is default
  // You can use `fallthrough` keyword
  switch 2 {
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two") // two
      // Can you fallthrough keyword if you want both lines to execute
      // fallthrough
    case 3, 4, 5:
      fmt.Println("three, four, five")
    default:
      fmt.Println("nothing found")
  }

  // Initializer syntax
  switch i := 2 + 3; i {
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    case 3, 4, 5:
      fmt.Println("three, four, five") // three, four, five
    default:
      fmt.Println("nothing found")
  }

  // Tagless syntax
  j := 10
  switch {
    case j <= 10:
      fmt.Println("less than 10")
      // Can you fallthrough keyword if you want both lines to execute
      fallthrough
    case j <= 20:
      fmt.Println("less than 20")
    default:
      fmt.Println("default")
  }

  // Type switch
  var k interface{} = 1
  switch k.(type) {
  case int:
    fmt.Println("type is: int") // type is: int
    if true {
      break
    }
    fmt.Println("this won't print since we broke out")
  case float64:
    fmt.Println("type is: float64")
  case string:
    fmt.Println("type is: string")
  default:
    fmt.Println("default")
  }
}

func returnTrue() bool {
  fmt.Println("Returning true...")
  return true
}
