package main

import (
  "fmt"
  "time"
  "sync"
  "runtime"
)

// WaitGroup
// Used to synchronization multiple goroutines
// Import "sync" package
// var wg = sync.WaitGroup{}
// Add `wg.Add(1)` above `go func`
// Call `wg.Done()` in your `go func`
// Add `sg.Wait()` below your `go func`
var wg = sync.WaitGroup{}
var counter = 0

// Mutex
// Basically a lock that app will honor
// Simple mutex is locked or unlocked
// When locked only one entity can access data at a time
//
// Rw Mutex
// Read / Write
// Unlimited read access
// Only one entity can obtain lock and write though
// Therefore, writer will wait until all the readers are done, then it will lock it,
// preventing reading and writing until it unlocks
var mut = sync.RWMutex{}
var counterMut = 0

func main() {

  // Goroutines

  // Creating goroutines
  // To turn into goroutine insert `go` keyword
  // Spin off a green thread and run sayHello function in green thread.
  // Most languages use OS threads, meanning they have an individual function 
  // call stack dedicated to the execution of whatever code is passed to that
  // thread.
  // Traditionally, these tend to be very large, about 1MB or ram, takes time to setup,
  // therefore, you end up with thread pools becuase creation and deletion of threads is expensive
  // We want to avoid this in languages like Java, C#

  // Green threads
  // Abstraction of thread called goroutine
  // Go runtime as a scheduler that will map these goroutines to the underlying OS threads
  // for periods of time.
  // We can now interact with higher level goroutines rather than threads
  // Therefore, goroutines can be created with vary little stack space becuase they
  // can be allocated quickly and destroyed quickly
  // Not uncommon for a Go program to handle thousands, or tens of thosands of 
  // goroutines at a time without a problem.

  go sayHello()

  // Don't do this!
  // Use sleep in order to get sayHello to execute
  // This is because the main function is itself a goroutine
  // and will exit immediately without any time for sayHello to execute
  time.Sleep(100 * time.Millisecond)

  closureWithRace()
  closure()

  // Synchronization

  // WaitGroups
  closureWG()

  for i := 0; i < 10; i++ {
    wg.Add(2)
    go sayHelloWG()
    go incrementWG()
  }
  wg.Wait()

  // Mutexes
  for i := 0; i < 10; i++ {
    wg.Add(2)
    // mut.RLock() // Putting lock outside of goroutine means
    // we are synchonously locking
    go sayHelloMut() // asynchronously unlocking in goroutine
    // mut.Lock() // Puttin lock outside of goroutine means we are now synchronous aka single threaded
    // This destroys multithreading, since it's not executing like a single thread
    // Therefore, if you needed to print things in the same order you can use this technique
    // or better yet, remove the wait group and mutex and execute normally as a single thread
    go incrementMut()
  }
  wg.Wait()

  // Parallelism
  // Max processes
  // By default number of threads will match number of cores
  // GOMAXPROCS is a tuning variable, minimum is one thread per OS core
  // However, you want to increase beyond number of cores to take advantage of parallelism.
  // Use performance testing with various values to find the optimal performance
  // Note -1 gives us number of threads previously set
  fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) // Threads: 8
  // runtime.GOMAXPROCS(1) // execute as single thread
  runtime.GOMAXPROCS(100)
  fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) // Threads: 100

  // Best practices
  // Don't create goroutines in libraries
  // Let consumer control concurrency
  // When creating a goroutine, know how it will end
  // Avoids subtle memory leaks
  // Check for race conditions at compile time use `-race` with `go build` then run app to see compiler's results
  // When using anonymous functions, pass data as local variables
}

func sayHello() {
  fmt.Println("Hello")
}

// Closure
// Generally considered a bad idea to acces via closure due to race condition
// Therefore, pass values into anonymous function
// Anonymous function has access to outer scope
// Race condition:
// - Most of the time, the scheduler will not interupt the main thread,
// therefore most of the time it will print "Goodbye" however it's not guaranteed
func closureWithRace() {
  var msg = "Hello"
  go func() {
    // Anonymous function has access to outerscope
    fmt.Println(msg)
  }()
  msg = "Goodbye"
  time.Sleep(100 * time.Millisecond)
}

func closure() {
  var msg = "Hello"
  go func(msg string) {
    // Pass msg into anonymous func to avoid race condition
    fmt.Println(msg)
  }(msg)
  msg = "Goodbye"
  time.Sleep(100 * time.Millisecond)
}

// WaitGroups
func closureWG() {
  var msg = "Hello"
  wg.Add(1)
  go func(msg string) {
    fmt.Println(msg)
    wg.Done()
  }(msg)
  msg = "Goodbye"
  time.Sleep(100 * time.Millisecond)
  wg.Wait()
}

func sayHelloWG() {
  fmt.Printf("Hello #%v\n", counter)
  wg.Done()
}

func incrementWG() {
  counter++
  wg.Done()
}

// Mutexes
func sayHelloMut() {
  // lock
  // action
  // unlock
  // done
  mut.RLock()
  fmt.Printf("Hello #%v\n", counterMut)
  mut.RUnlock()
  wg.Done()
}

func incrementMut() {
  // lock
  // action
  // unlock
  // done
  mut.Lock()
  counterMut++
  mut.Unlock()
  wg.Done()
}
