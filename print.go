package main

import "fmt"

func main() {
  fmt.Println("Let's first see how", "the Println() method works.")
  fmt.Println("Notice that each statement adds a newline for us.")
  fmt.Println("There's also a default space", "between the string arguments.")
  
  // Does not automatically add spaces or line breaks, like Println does
  fmt.Print("Print", "is", "different")
  fmt.Print("See?")
}
