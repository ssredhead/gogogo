package main

import "fmt"

func main() {
	name := "H. J. Simp"

	switch name {
    case "Butch": 
      fmt.Print("Head to Robbers Roost.")
    case "Bonnie":
      fmt.Print("Stay put in Joplin.")
    default:
      fmt.Print("Just hide!")   
  }
}
