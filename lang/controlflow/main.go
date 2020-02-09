package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func main() {

  // Looping

  // for
  // break keyword
  // continue keyword
  fmt.Println("\nfor loop:")

  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }

  // Multiple variables
  // Scoped to for loop
  for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
    fmt.Println(i)
    fmt.Println(j)
  }

  // Split declaration
  i := 0 // i now scoped to main function
  for ; i < 5; i++ {
    fmt.Println(i)
  }

  // While loop ~ sort of
  fmt.Println("\npseudo while loop:")
  j := 0
  for j < 5 {
    fmt.Println(j)
    j++
  }

  b := true
  for b {
    fmt.Println(b)
    b = false
  }

  for true {
    fmt.Println(b)
    if true {
      break
    }
  }

  // Range keyword
  // Use to loop over collections (arrays, slices, maps)
  fmt.Println("\nRange:")
  s := []int{1, 2, 3}
  for k, v := range s {
    fmt.Println(k, v)
    // 0 1
    // 1 2
    // 2 3
  }

  ar := [...]int{1, 2, 3}
  for k, v := range ar {
    fmt.Println(k, v)
    // 0 1
    // 1 2
    // 2 3
  }

  m := map[string]int{
    "key1": 1,
    "key2": 2,
  }
  for k, v := range m {
    fmt.Println(k, v)
    // key1 1
    // key2 2
  }

  str := "hello"
  for k, v := range str {
    // v will be ASCII number, covert to string
    fmt.Println(k, v, string(v))
    // 0 104 h
    // 1 101 e
    // 2 108 l
    // 3 108 l
    // 4 111 o
  }

  // Only keys
  for k := range m {
    fmt.Println(k)
    // key1
    // key2
  }

  // Only values
  for _, v := range m {
    fmt.Println(v)
    // 1
    // 2
  }

  // Channels
  // used for concurrency
  // can range over them...

  // Defer
  // Envoke function but delay execution
  // Executes any defered functions that were passed into it after the outer function exits, but before any values are returned
  // Runs in LIFO order
  // Commonly used to close resources
  // Takes a function all as input, can be an anoymous function, but anonymous function must be called.
  fmt.Println("\nDefer:")
  defered := deferFn()
  fmt.Println(defered) // 22
  closure()

  // Panic
  // Go returns errors rather than exceptions for common exceptions in other languages,
  // like trying to open a file that doesn't exist
  // Only use when the program is in a state the can't be recovered from e.g. web server can't start
  // Panic happens after deferred statements,
  // which makes sense since you want to close resource before exiting program
  fmt.Println("\nPanic:")

  // If you uncomment the code below the program will exit after the panic
  // panicEx()

  // Recover
  // When you run into panic, how to recover
  // Only usefule in deferred functions
  // Current function will not attempt to continue, but higher functions in call stack will
  fmt.Println("\nRecover:")

  // pre panicker
  // about to panic
  // this anonymous function will immediately envoke
  // 2020/02/07 18:54:04 Error:  uh oh
  // post panicker
  fmt.Println("pre panicker")
  panicker()
  fmt.Println("post panicker")
}

// Will print:
// deferFn...
// start
// end
// <deferFn has finished at this point>
// Now runtime sees it has a defered function to call
// runtime calls that function then returns value
// middle
func deferFn() int {
  fmt.Println("deferFn...")
  fmt.Println("start")
  defer fmt.Println("middle")
  fmt.Println("end")
  return 22
}

func closure() {
  a := "Start"
  defer fmt.Println(a)
  a = "end"
}

func httpWithDefer() {
  res, err := http.Get("http://www.google.com/robots.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  robots, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s", robots)
}

func panicEx() {
  n1, n2 := 1, 0
  if n1 == 0 || n2 == 0 {
    panic("I can't divide by zero!") // apparently, automatically prints stack trace too
  } else {
    ans := n1 / n2 // 1 / 0 error
    fmt.Println(ans)
  }
}

// Automatically envoke anonymous function with braces after declaration
// Will print
// start
// Error: uh oh
func panicker() {
  fmt.Println("about to panic")
  func() {
    fmt.Println("this anonymous function will immediately envoke")
  }()
  defer func() {
    if err := recover(); err != nil {
      log.Println("Error: ", err)
      // If you can't handle the error then rethrow the panic
      // panic(err)
    }
  }()
  panic("uh oh")
  fmt.Println("done panicking") // this will never print
}
