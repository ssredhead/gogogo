package main

import "fmt"

func main() {
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"

  // comparison operators:
  if lockCombo == robAttempt {
    fmt.Print("The vault is now opened.")
  }
  
}
