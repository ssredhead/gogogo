package main

import "fmt"

// Make brainwash have a pointer parameter
// using them all at once
func brainwash(saying *string) {

	// Dereference saying: 
	*saying = "Beep Boop."
}

func main() {
	greeting := "Hello there!"
	
	// Call by reference
	brainwash(&greeting)
	
	fmt.Println("greeting is now:", greeting)
}
