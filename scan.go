package main

import "fmt"

func main() {
  fmt.Println("What would you like for lunch?")
  
  // Getting input with scan:
  var food string

  fmt.Scan(&food)
  
  fmt.Printf("Sure, we can have %v for lunch.", food)
}
