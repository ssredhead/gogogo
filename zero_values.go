package main

import "fmt"

func main() {
  // Create three variables
  //By default, numbers are given a value of 0 if not initialized, strings are given "" and bools are given false
  // emptyInt an int8
  var emptyInt int8
  // emptyFloat a float32
  var emptyFloat float32
  // and emptyString a string
  var emptyString string
  // Finally, print them all out
  fmt.Println(emptyInt, emptyFloat, emptyString)
}