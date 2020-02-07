package main

import (
  "fmt"
)

func main() {

  // Interfaces
  // Are types
  // Naming convention: Single method should be name of interface + "er"
  // e.g. Write -> Writer, Read -> Reader
  // Name it by what it does

  // Struct is most common way to implement an interface.
  // However, you can use any type that can have a method associated to it
  // can implement an interace.
  // Furethermore any type can have a methods associated to it
  var w Writer = ConsoleWriter{}
  w.Write([]byte("Hello Go!"))

  // Type implementing interface
  myInt := IntCounter(0) // casting int to IntCounter type alias
  var inc Incrementer = &myInt
  for i := 0; i < 10; i++ {
    fmt.Println(inc.Increment())
  }

  // Composing interfaces

  // Type conversion

  // Empty interface
  // Interface defined on the fly with no methods on it
  // Every type in Go implements the empty interface
  // Everything can cast into an object that has no methods on it, including primitives
  // However, you must:
  // - Cast to an object (type conversion)
  // - Use "reflect" package to find out what methods are availble on the object
  var myObj interface{}
  fmt.Println(myObj) // <nil>

  // Type switches
  var inf interface{} = 0
  switch inf.(type) {
    case int:
      fmt.Println("It's an integer") // it's an integer
    case string:
      fmt.Println("It's a string")
    default:
      fmt.Println("I don't know what it is")
  }

  // Implementing with values vs. pointers
  // Method set of value is all methods with value receivers
  // Method set of pointer is all methods, regardless of reciever type
  // i.e. all of the value recievers and all of the pointer receivers
  // pointer types are more flexible, however it does allow for consumer to modify
  // the data of the object your are passing in.

  // Best practices
  // Use many small interfaces
  // Single method interfaces are some of the most powerful and flexible:
  //  - io.Writer, io.Reader, interface{}
  // 
  // Compose small interfaces rather than creating a large interface 
  //
  // Types that will be consumed
  // Don't export interfaces for types that will be consumed
  // Publish the concrete type, rather than publishing an interface
  // Allow consumers to define the methods they need, otherwise, the consumer 
  // will be force to implement a bunch of methods they may not need or use
  // 
  // Do export interfaces for types that you will be consuming in your package
  // That way users of your package can create the concrete implementation you need.
  // That way you don't need to worry about the implementation, only the behvaior that are being exposed.
  // 
  // This is contrary to most languages due to Go preferring implicit implementation vs explicit implementation
  // Library:
  // - Accept interfaces
  // - Export reasonable defaults
  // - Consumers of your library can implement the interface however they need
  // Design functions and methods to receive interfaces whenever possible
  // Compose interfaces together
  // e.g. Writer, Closer, WriterCloser, etc
}

// Struct implementing interface
//------------------------------------------------------------------------------ 
// Interfaces describe behavior
type Writer interface {
  Write([]byte) (int, error)
}

// Interfaces are implicit rather than explicit
type ConsoleWriter struct {}

// Interfaces are implicit rather than explicit
func (cw ConsoleWriter) Write(date []byte) (int, error) {
  n, err := fmt.Println(string(date))
  return n, err
}

// Type implementing interface
//------------------------------------------------------------------------------ 
type Incrementer interface {
  Increment() int
}

// Type alias for int
// int is a primitive so you can't modify it
// you can however create a type alias to int and modify it
// int is storing the data, that the method is using
type IntCounter int

func (ic *IntCounter) Increment() int {
  *ic++
  return int(*ic)
}
