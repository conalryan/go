package main

import (
  "fmt"
  "sync"
  "time"
)

var wg = sync.WaitGroup{}

const (
  logInfo = "INFO"
  logWarning = "WARNING"
  logError = "ERROR"
)

type logEntry struct {
  time time.Time
  severity string
  message string
}

var logCh = make(chan logEntry, 50)
// Signal only channel
var doneCh = make(chan struct{})

func main() {

  // Channels
  // Most languages were built with a single processing core in mind
  // Therefore, parallelism and concurrency was added in
  // However, Go was born in mutli-core processing world
  // goroutines are an abstraction on top of threads
  // Channels allow the safe passing/sharring of data between goroutines
  // Designed to synchronize data transmission between goroutines

  // Channel basics
  // Create using make function
  // Use `chan` keyword in `make` function and specify data type that will flow through channel
  // Strongly typed
  ch := make(chan int)
  wg.Add(2)

  // Sending and receiving
  // Arrow points to where you want data to flow
  // You can have different number of senders and receivers

  // Receiver
  // Restricting data flow:
  // Instead of go func(), make explicit receive only channel
  go func(ch <-chan int) {
    // Receive the value from the channel
    i := <- ch
    fmt.Println(i)
    wg.Done()
  }(ch)

  // Sender
  // Restricting data flow:
  // Instead of go func(), make explicit send only channel
  go func(ch chan<- int) {
    i := 42
    // Send 42 into the channel
    // Passed by value, go makes a copy of the value, so change below has no affect
    ch <- i
    i = 27
    wg.Done()
  }(ch)

  wg.Wait()

  // Buffered channels
  // By default Channels block sender side til receiver is available
  // By default Block receiver side til message is available
  // To decouple this, use buffered channel 
  // Specify 2nd parameter to create internal store.
  // e.g. chan int, 50 will store 50 integers
  bufCh := make(chan int, 50)

  var bufWg = sync.WaitGroup{}
  bufWg.Add(2)

  // Receiver single receiver looping over mulitple messages
  go func(ch <-chan int) {
    // Loop over range of messages in channel
    for i := range ch {
      fmt.Println(i)
    }
    bufWg.Done()
  }(bufCh)

  // Sender sending multiple values
  go func(ch chan<- int) {
    // Send multiple values
    ch <- 42
    ch <- 27
    // Close channnel when you're done sending values to prevent dead lock
    close(ch)
    bufWg.Done()
  }(bufCh)

  bufWg.Wait()

  // Select statements
  // Similar to switch statements but only used with channels
  // Allows goroutine to monitor severl channels at once
  // Blocks if all channels block
  // If multiple channesl receive value simulataneously, behavior is undefined
  // i.e. the order of the select statments doesn't matter
  // Can be blocking (no default case specificed) or non-blocking (default case specified)
  // This is because the channles block by default, but the default case can execute
  go logger()
  logCh <- logEntry{time.Now(), logInfo, "App is starting"}

  logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
  time.Sleep(100 * time.Millisecond)
  doneCh <- struct{}{} // struct{} is type, second curly braces initializes the empty struct
}

func logger() {
  for {
    // Blocking select statement
    // Note if you add a default case then it becomes a non-blocking select statement
    select {
    case entry := <-logCh:
      fmt.Printf("%v - [%v]%v\n", entry.time.Format("2020-02-08T10:43:00"), entry.severity, entry.message)
    case <- doneCh:
      break
    }
  }
}
