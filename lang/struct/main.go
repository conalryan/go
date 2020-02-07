package main

import (
  "fmt"
  "reflect"
)

// internal lowerCase
type internalStruct struct {
  someInt int
  somebool bool
}

// exported UpperCase
type ExportedStruct struct {
  SomeInt int
  SomeBool bool
}

type Person struct {
  Age int
  Name string
  Friends []string
}

// Embedding
// go does not support inheritance
// only supports composition
type Animal struct {
  Name string
  Origin string
}

// Tags
// Use backticks
// go uses reflection to get values of tags
// import "reflect"
type TagStruct struct {
  Name string `required max: "100"`
}

type Bird struct {
  Animal // Embeds Animal...Bird "has" animal characteristics
  CanFly bool
}

func main() {

  // Structs
  fmt.Println("\nStructs:")
  // Struct is a value type
  // Do not use positional sytax when you declare a struct,
  // always specify the field names. The order of the fields does not matter.
  // Capital letter is exported the struct name and the fields
  // lowercase letter is internal to package

  // Literal syntax
  a := Person{
    Age: 22,
    Name: "Bob",
    Friends: []string{
      "George",
      "Sally",
    },
  }
  fmt.Printf("%v, %T\n", a, a) // {22 Bob [George Sally]}, main.Person

  // Dot syntax
  fmt.Printf("%v, %T\n", a.Name, a.Name) // Bob, string

  // Anonymous struct
  b := struct{name string}{name: "Frank"}
  fmt.Printf("%v, %T\n", b, b) // {Frank}, struct { name string }

  // Copy struct
  c := a
  c.Name = "Robert"
  fmt.Printf("%v, %T\n", c, c) // {22 Robert [George Sally]}, main.Person
  fmt.Printf("%v, %T\n", a, a) // {22 Bob [George Sally]}, main.Person

  // Address of operator &
  d := &a
  d.Age = 44
  fmt.Printf("%v, %T\n", d, d) // &{44 Bob [George Sally]}, *main.Person
  fmt.Printf("%v, %T\n", a, a) // {44 Bob [George Sally]}, main.Person

  // Embedding
  e := Bird{
    Animal: Animal{Name: "Emu", Origin: "Australia"},
    CanFly: false,
  }
  fmt.Printf("%v, %T\n", e, e) // {{Emu Australia} false}, main.Bird

  // Tags
  t := reflect.TypeOf(TagStruct{})
  field, _ := t.FieldByName("Name")
  fmt.Println(field.Tag) // required max: "100"
}
