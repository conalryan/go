package main

import (
  "fmt"
)

func main() {

  // bool
  fmt.Println("\nbool:")
  // true or false
  // not a type alias, own type
  // defaults to false 
  var defaultBool bool
  fmt.Printf("%v, %T\n", defaultBool, defaultBool) // false, bool

  var aBool bool = false
  fmt.Printf("%v, %T\n", aBool, aBool) // false, bool

  // Result of expression
  result := 1 == 1
  fmt.Printf("%v, %T\n", result, result) // true, bool

  // Numeric

  // int
  fmt.Println("\nint:")
  // `int` Default e.g. `n := 42`
  // `int8` -128 0 127
  // `int16` -32,768 - 32,767
  // `int32` -2.147 billion - 2.147 billion
  // `int64` -9 quintillion - 9 quintillion
  // big package from math library for larger
  var defaultInt int
  fmt.Printf("%v, %T\n", defaultInt, defaultInt) // 0, int

  i := 42
  fmt.Printf("%v, %T\n", i, i) // 42, int

  // unsigned
  fmt.Println("\nunsinged int:")
  // uint8 0 - 255
  // uint16 0 - 65,535
  // uint32 0 - 4.294 billion
  var ui uint16 = 56
  fmt.Printf("%v, %T\n", ui, ui) // 56 uint16

  // byte
  fmt.Println("\nbyte:")
  // alias for uint8
  var defaultByte byte
  fmt.Printf("%v, %T\n", defaultByte, defaultByte) // 0, uint8

  a := 10 // 1010
  b := 3 // 0011
  fmt.Println(a & b) // 0010
  fmt.Println(a | b) // 1011 
  fmt.Println(a ^ b) // 1001
  fmt.Println(a &^ b) // 0100

  // bit shifting
  fmt.Println("\nbit shifting:")
  a = 1
  fmt.Println(a << 1) // 2 
  fmt.Println(a >> 1) // 0

  // floating point
  fmt.Println("\nfloating point numbers:")
  // `float32` 
  // `float64`
  var defaultFloat float32
  fmt.Printf("%v, %T\n", defaultFloat, defaultFloat) // 0, float32

  var f32 float32 = 3.14
  var f64 float64 = 2.1E14
  fmt.Printf("%v, %T\n", f32, f32) // 3.14, float32
  fmt.Printf("%v, %T\n", f64, f64) // 2.1e+14, float64

  // complex
  fmt.Println("\ncomplex numbers:")
  // great for data science
  // complex64 combines float32 and float32 for real and imaginary parts
  // complex128 combines float64 and float64 for real and imaginary parts
  var defaulatComplex complex64
  fmt.Printf("%v, %T\n", defaulatComplex, defaulatComplex) // (0+0i), complex64

  var c64 complex64 = 1 + 2i
  var c128 complex128 = 3i
  fmt.Printf("%v, %T\n", c64, c64) // (1+2i), complex64
  fmt.Printf("%v, %T\n", c128, c128) // (0+3i), complex128

  // extract real and imaginary parts
  fmt.Printf("%v, %T\n", real(c64), real(c64)) // 1, float32
  fmt.Printf("%v, %T\n", imag(c64), imag(c64)) // 2, float32

  var b64 complex64 = 2 + 5.2i
  fmt.Println(c64 + b64) // (3+7.2i)
  fmt.Println(c64 - b64) // (-1-3.1999998i)
  fmt.Println(c64 * b64) // (-8.4+9.2i)
  fmt.Println(c64 / b64) // (0.39948454-0.03865979i)

  // complex function
  // creates a complex number
  var cComplex complex128 = complex(5, 12)
  fmt.Printf("%v, %T\n", cComplex, cComplex) // (5+12i), complex128

  // text types

  // string
  fmt.Println("\nstring:")
  // any UTF8 character
  // alias for byte
  // immutable
  // can be concatenated with + operator
  var defaultStr string
  fmt.Printf("%v, %T\n", defaultStr, defaultStr) // , string

  s := "this is a string"
  fmt.Printf("%v, %T\n", s, s) // this is a string, string

  // Use brackets[] to access chars
  fmt.Printf("%v, %T\n", s[2], s[2]) // 105, unit8

  var s2 string = "append me"
  fmt.Printf("%v, %T\n", s + s2, s + s2) // this is a stringappend me, string

  // convert to slice of bytes
  sB := []byte(s)
  fmt.Printf("%v, %T\n", sB, sB) // [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8

  // rune
  fmt.Println("\nrune:")
  // any UTF32 char
  // declared using single quotes
  // type alias for int32
  var defaultRune rune
  fmt.Printf("%v, %T\n", defaultRune, defaultRune) // 0, int32

  r := 'a'
  fmt.Printf("%v, %T\n", r, r) // 97, int32

  var aRune rune = 'b'
  fmt.Printf("%v, %T\n", aRune, aRune) // 98, int32
}
