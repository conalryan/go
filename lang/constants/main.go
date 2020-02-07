package main

import (
  "fmt"
)

const aConst int16 = 27

// iota
// counter you can use for enumerated constants
const (
  a = iota
  b = iota
  c = iota
)

// Compiler infers pattern and applies "= iota" to e and f
const (
  d = iota
  e
  f
)

// Best Practice
// Use "Zero" value as error value
// Underscore _ is go's write only variable (throw away value)
// c.r. similar to TypeScript using 0 as base value in enum,
// you can get a false positive since the default value can be 0,
// and iota will initialize to 0
const (
  _ = iota
  g
  h
)

const (
  _ = iota // ignore first value by assigning to blank identifier
  KB = 1 << (10 * iota)
  MB
  GB
  TB
  PB
  EB
  ZB
  YB
)

const (
  isAdmin = 1 << iota
  isCorporate
  isRegional

  canView
  canManage
)

func main() {

  // Constants
  fmt.Println("\nConstants")
  // const keyword
  // Immutable
  // Use camelCase for non-exported constants
  // Use PascalCase for exported constants
  // Value must be assigned at compile type
  // Can be shadowed - inner most scope wins
  const aNonExportedConstant string = "Package private constant"
  fmt.Printf("%v, %T\n", aNonExportedConstant, aNonExportedConstant) // Package private constant, string

  const AnExportedConstant string = "Exported constant"
  fmt.Printf("%v, %T\n", AnExportedConstant, AnExportedConstant) // Exported constant, string

  // Untype constant
  const n42 = 42
  fmt.Printf("%v, %T\n", n42, n42) // 42, int
  // You can now use the constant with any number type int8, int32, etc
  var fn42 int8 = n42
  fmt.Printf("%v, %T\n", fn42, fn42) // 42, int8

  // Shadowing
  const aConst int = 14
  fmt.Printf("%v, %T\n", aConst, aConst) // 14, int

  // Enumerated Constants
  fmt.Println("\nEnumerated Constatns:")
  fmt.Printf("%v, %T\n", a, a) // 0, int
  fmt.Printf("%v, %T\n", b, b) // 1, int
  fmt.Printf("%v, %T\n", c, c) // 2, int
  fmt.Printf("%v, %T\n", d, d) // 0, int
  fmt.Printf("%v, %T\n", e, e) // 1, int
  fmt.Printf("%v, %T\n", f, f) // 2, int

  fmt.Printf("%v, %T\n", g, g) // 1, int
  fmt.Printf("%v, %T\n", h, h) // 2, int

  // File size with bit shifting
  fileSize := 4000000000.
  fmt.Printf("%.2fGB\n", fileSize/GB) // 3.73GB

  // BitMask roles
  fmt.Printf("%b\n", isAdmin) // 1
  fmt.Printf("%b\n", isCorporate) // 10
  fmt.Printf("%b\n", isRegional) // 100

  var canAccess byte = isAdmin | isCorporate | canManage
  fmt.Printf("%b\n", canAccess) // 10011

  var someUser byte = isCorporate | canManage | canView
  fmt.Printf("Is admin? %v\n", isAdmin & someUser == isAdmin) // Is admin? false
  fmt.Printf("Is corporate? %v\n", isCorporate & someUser == isCorporate) // Is corporate? true
}
