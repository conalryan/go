package main

import (
  "fmt"
)

func main() {

  fmt.Println("\nCollections:")

  // Array
  fmt.Println("\nArrays:")
  // Arrays are considered values, rather than pointer to array
  // Therefore, when you copy an array, go will make a new copy of the array
  // a := [3]int{1, 2, 3}
  // a := [...]int{1, 2, 3}
  // a := [3]int
  var defaultArr [3]int // can hold 3 numbers
  fmt.Printf("%v, %T\n", defaultArr, defaultArr) // [0 0 0], [3]int

  defaultArr[1] = 22
  fmt.Printf("%v, %T\n", defaultArr, defaultArr) // [0 22 0], [3]int

  // initializer syntax
  strArr := [...]string{"hi", "there"}
  fmt.Printf("%v, %T\n", strArr, strArr) // [hi there], [2]string
  fmt.Printf("%v, %T\n", strArr[0], strArr[0]) // hi, string

  // Length function
  fmt.Printf("%v\n", len(strArr)) // 2

  // Copy array - copy values
  a := [...]int{1, 2, 3}
  b := a
  b[1] = 5
  fmt.Printf("%v, %T\n", a, a) // [1 2 3], [3]int
  fmt.Printf("%v, %T\n", b, b) // [1 5 3], [3]int

  // Copy array reference using reference operator &
  c := &a
  c[1] = 5
  fmt.Printf("%v, %T\n", a, a) // [1 5 3], [3]int
  fmt.Printf("%v, %T\n", c, c) // [&[1 5 3], *[3]int

  // Slice
  fmt.Println("\nSlices:")
  // Reference type
  // Initialized as literal with square brackets
  // Slicing operations are available on slices and arrays
  // Backed by array
  var defaultSlice []int
  fmt.Printf("%v, %T\n", defaultSlice, defaultSlice) // [], []int

  sa := []int{1, 2, 3, 4, 5, 6, 7}
  sb := sa[:] // slice of all elements
  sc := sa[2:] // slice from 3rd element to end
  sd := sa[:5] // slice of first 5 elements
  se := sa[3:6] // slice of 4th, 5th and 6th elements
  fmt.Printf("%v, %T\n", sa, sa) // [1 2 3 4 5 6 7], []int
  fmt.Printf("%v, %T\n", sb, sb) // [1 2 3 4 5 6 7], []int
  fmt.Printf("%v, %T\n", sc, sc) // [3 4 5 6 7], []int
  fmt.Printf("%v, %T\n", sd, sd) // [1 2 3 4 5], []int
  fmt.Printf("%v, %T\n", se, se) // [4 5 6], []int

  // Length of slice
  fmt.Printf("%v\n", len(sa)) // 7 

  // Capacity of slice
  fmt.Printf("%v\n", cap(sa)) // 7

  // Copy slices - copy reference
  s2 := sa
  s2[1] = 5 // This will change sa too, and all slices derived from sa
  fmt.Printf("%v, %T\n", sa, sa) // [1 5 3 4 5 6 7], []int
  fmt.Printf("%v, %T\n", s2, s2) // [1 5 3 4 5 6 7], []int
  fmt.Printf("%v, %T\n", sd, sd) // [1 5 3 4 5], []int

  // Make function
  fmt.Println("\nmake function:")
  m := make([]int, 3)
  fmt.Printf("%v, %T\n", m, m) // [0 0 0], []int
  fmt.Printf("%v\n", len(m)) // 3
  fmt.Printf("%v\n", cap(m)) // 3

  // Set capacity
  mc := make([]int, 3, 100)
  fmt.Printf("%v, %T\n", mc, mc) // [0 0 0], []int
  fmt.Printf("%v\n", len(mc)) // 3
  fmt.Printf("%v\n", cap(mc)) // 100

  // Append function
  // go usually doubles size of array when appending
  mi := []int{} // empty slice len and cap of 0
  mi = append(mi, 1)
  fmt.Printf("%v, %T\n", mi, mi) // [1], []int
  fmt.Printf("%v\n", len(mi)) // 1
  fmt.Printf("%v\n", cap(mi)) // 1

  // Append multiple values
  mi = append(mi, 2, 3, 4)
  fmt.Printf("%v, %T\n", mi, mi) // [1 2 3 4], []int
  fmt.Printf("%v\n", len(mi)) // 4
  fmt.Printf("%v\n", cap(mi)) // 4

  sf1 := []int{1, 2, 3, 4, 5}
  fmt.Printf("%v\n", sf1) // [1, 2, 3, 4, 5]

  sf2 := sf1[1:] // shift i.e. remove first element
  fmt.Printf("%v\n", sf2) // [2, 3, 4, 5]

  sf3 := sf1[:len(sf1)-1] // remove last element
  fmt.Printf("%v\n", sf3) // [1, 2, 3, 4]

  sf4 := append(sf1[:2], sf1[3:]...) // remove middle
  fmt.Printf("%v\n", sf4) // [1, 2, 4, 5]

  // Maps
  fmt.Println("\nMaps:")
  // Key type must be able to be tested for equality
  // strings, bool, numeric, arrays, pointers
  // Slices, maps and other functions cannot be used as a key.
  // Return order of a map is not guaranteed
  defaultMap := map[int]string{}
  fmt.Printf("%v, %T\n", defaultMap, defaultMap) // map[], map[int]string

  // Literal syntax
  aMap := map[string]int{
    "key1": 22,
    "anotherKey": 99,
  }
  fmt.Println(aMap) // map[anotherKey:99 key1:22]

  // Make function
  mMap := make(map[string]int)
  fmt.Println(mMap) // map[]

  // Bracket syntax
  // retrieve value from map
  fmt.Println(aMap["key1"])

  // edit value in map
  aMap["anotherKey"] = 100
  fmt.Println(aMap) // map[anotherKey:100 key1:22]

  // add new key/value to map
  key := "brandNewKey"
  aMap[key] = 11
  fmt.Println(aMap) // map[anotherKey:100 brandNewKey:11 key1:22]

  // Delete function
  // remove key/value from map
  delete(aMap, key)
  fmt.Println(aMap) // map[anotherKey:100 key1:22]

  // Comma Ok syntax
  value, ok := aMap["key1"]
  if ok {
    fmt.Printf("%v, %T\n", value, value) // 22, int
  }

  if _, ok := aMap["key1"]; ok {
    fmt.Println("Ya the key is there")
  }

  // Length function
  fmt.Printf("%v\n", len(aMap)) // 2

  // Copy map
  cMap := aMap
  cMap["oops"] = 45
  fmt.Println(cMap) // map[anotherKey:100 key1:22 oops:45]
  fmt.Println(aMap) // map[anotherKey:100 key1:22 oops:45]
}
