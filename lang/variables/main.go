package main

import (
  "fmt"
  "strconv"
)

// cd <to this dir>
// go build
// ./variables

// Package level variable declaration
// Can't use := syntax
var pkgInt int = 32

// Var block
// Syntax sugar to avoid multiple lines with 'var'
var (
  aString string = "some string"
  afloat float32 = 22.32
)

// Can have multiple var blocks
var (
  featA string = "feature a stuff"
  id int = 22
)

func main() {

  fmt.Println("\nVariable declarations:")
  // Option 1: Declare variable
  var i int
  // Assign value to variable
  i = 42
  fmt.Printf("%v, %T\n", i, i)

  // Option 2: Declare variable and assign value
  var j int = 42
  fmt.Printf("%v, %T\n", j, j)

  // Option 3: Syntax sugar - declare variable and assing value 
  // Can only use this syntax in func block
  // Can't specify certain types such as float32, go will default to float64 with decimal number
  k := 42
  fmt.Printf("%v, %T\n", k, k)

  // Variable declared but not used is compile error
  // ./main.go:50:3: x declared and not used
  // x := "won't compile"

  // Shadowing
  fmt.Println("\nShadowing:")
  fmt.Printf("%s\n", aString)
  aString := "another string"
  fmt.Printf("%s\n", aString)

  // Type convertions
  fmt.Println("\nType convertions:")
  var anInt int = 32
  var aFloat float32
  aFloat = float32(anInt)
  fmt.Printf("%v, %T\n", aFloat, aFloat)

  var typeInt int = 2
  var typeStr string
  typeStr = strconv.Itoa(typeInt)
  fmt.Printf("%v, %T\n", typeInt, typeStr)
}

