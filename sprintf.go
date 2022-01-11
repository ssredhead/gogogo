package main

import "fmt"

func main() {
  template := "I wish I had a %v."
  pet := "puppy"
  
  var wish string
  
  // formatting + interpolation:
  wish = fmt.Sprintf(template, pet)
  
  
  fmt.Println(wish)
}
