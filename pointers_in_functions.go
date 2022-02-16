package main

import "fmt"

// Make brainwash have a pointer parameter
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
