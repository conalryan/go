[golang](https://golang.org/)
================================================================================

[Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=3456s)
================================================================================

History
--------------------------------------------------------------------------------
- Created by Google by:
  - Robert Griesemar
  - Rob Pike
  - Ken Thompson
    - Build fist Unix systen and helped develop Unicode
- Looking to solve issues with primary languages used at Google:
  - Java:
    - Is quick but type system continues to become more complex over time (natural trend of language)
    - Complex type system
    - Code base continues to bloat over time
    - Originally built with single thread system
    - Concurrency was patched in at best
  - Python:
    - Interpreted language
    - Slow run time, at scale
    - Only supports single thread
  - C/C++:
    - Complex type system
    - Build times are notorously slow
    - Concurrency was patched in at best

Go solutions
--------------------------------------------------------------------------------
- Strongly and statically typed:
  - Can't change data type
  - Variable have to defined at compile time
- Focus on simplicity
- Focus on fast compile times
- Garbage collected language (for simplicity)
  - You can manage memory, however it's not necessary (cr. Really, how so?)
- Concurrency is built in
- Compiles down to a single binary

Workspace
--------------------------------------------------------------------------------
`src` All you code lives inside a single src directory

`bin` Final compiled binary 

`pkg` Intermediate binaries are placed here

Package
--------------------------------------------------------------------------------
- Every go application is strucuted into packages
- Every go file will need to declare what package it's a part of
- `main` is a special package because it is the entry point into your application
- Names should be short consice, typically nouns, all lower case

