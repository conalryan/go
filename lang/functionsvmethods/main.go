package main

import (
  "fmt"
)

func main() {

  // Execution syntax

  // Functions
  // - Called independently with the arguments specified
  np := NewPerson("John", 21)
  fmt.Println(np.isAdult()) // true

  // Methods
  // = Called on the type of their receiver:
  p := &Person{}
	p = p.withName("John").withAge(21)
  fmt.Println(*p) // {John 21}
}

type Person struct {
  Name string
  Age int
}

// Syntax

// Declaration syntax

// Functions
// Declared by specifying:
// - Types of the arguments
// - Return values
// - Function body
// Use for stateless actions
// “Stateless” here means any piece of code that always returns the same output for the same input
//
// This function returns a new instance of `Person`
func NewPerson(name string, age int) *Person {
  return &Person{
    Name: name,
    Age: age,
  }
}

// Method
// A method is a function that has a defined receiver, in OOP terms, a method is a function on an instance of an object.
// Go does not have classes. However, you can define methods on struct types.
// The method receiver appears in its own argument list between the func keyword and the method name.
//
// In fact, you can define a method on any type you define in your package, not just structs. You cannot define a method on a type from another package, or on a basic type.
// To define methods on a type you don’t “own”, you need to define an alias for the type you want to extend:
//
// Remember that Go passes everything by value, meaning that when Greeting() is defined on the value type, every time you call Greeting(), you are copying the User struct. Instead when using a pointer, only the pointer is copied (cheap).
//
// The other reason why you might want to use a pointer is so that the method can modify the value that its receiver points to.
//
// A method on the other hand is declared by additionally specifying the “receiver”
// (which in OOP terms would be the “class” that the method is a part of):
func (p *Person) isAdult() bool {
  return p.Age > 18
}
// In the above method declarations, we declared the isAdult method on the *Person type.

// Method chaining
// One very useful property of methods is the ability to chain them together, while still keeping your code clean. Let’s take an example of setting some attributes of Person using chaining:
func (p *Person) withName(name string) *Person {
	p.Name = name
	return p
}

func (p *Person) withAge(age int) *Person {
	p.Age = age
	return p
}

// Factory pattern
// Use object to create new instance of an object
// Anti pattern in Go and should be avoided

// The corollary here, is that if you find that a function reads and modifies a lot of values of a particular type, it should probably be a method of that type.

// Semantics
// Semantics refers to how the code read. If you read the code aloud in your spoken language, what makes more sense?
