package main

import (
  "fmt"
)

func main() {

  // Creating pointers
  fmt.Println("\nPointers:")

  // Default assignment is copy by value
  ca := 22
  cb := ca
  fmt.Println(ca, cb) // 22, 22 
  cb = 77
  fmt.Println(ca, cb) // 22, 77

  // Create pointer
  // Address of operator &
  var a int = 22
  var b *int = &a // b now holds the memory location of a
  fmt.Println(&a, b) // 0xc00009a038 0xc00009a038

  // Dereferencing pointers
  // Dereferencing operator asterisk *
  fmt.Println(a, *b) // 22 22

  // Modifying value
  a = 27
  fmt.Println(a, *b) // 27 27
  *b = 14
  fmt.Println(a, *b) // 14 14

  // Pointer arithmetic
  ara := [3]int{1, 2, 3}
  arb := &ara[0]
  arc := &ara[1]
  // Note memory locations are 8 bytes apart
  fmt.Printf("%v %p %p\n", ara, arb, arc) // [1 2 3] 0xc000018260 0xc000018268

  // However, you cannot do pointer arithmetic
  // Error: ./main.go:43:18: invalid operation: &ara[1] - 4 (mismatched types *int and int)
  // ard := &ara[1] - 8

  // Pointer types
  // Working with nil
  // If your function accepts pointers, you must check that the point isn't nil
  var ms *myStruct
  fmt.Println(ms) // <nil>
  ms = &myStruct{foo: 42}
  fmt.Println(ms) // &{42}

  // The new function
  // You can't initialize any values though
  var ms2 *myStruct
  fmt.Println(ms2) // <nil>
  ms2 = new (myStruct)
  fmt.Println(ms2) // &{0}

  // Set value - dereference pointer
  (*ms2).foo = 22
  fmt.Println((*ms2).foo) // 22

  // Syntax sugar
  // Compiler understands we really want the value of the pointer not the address
  ms2.foo = 99
  fmt.Println(ms2.foo) // 99 

  // Types with internal pointers
  // Slices are pointers to the first element of the underlying array
  // Maps are pointers to underlying data
  // - All assignemtn operations in Go are copy operations
  // - Slices and maps contain internal pointer, so copies point to same underlying data
}

type myStruct struct {
  foo int
}
