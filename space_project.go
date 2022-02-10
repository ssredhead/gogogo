package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("You have", fuel, "gallons of fuel left!")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
    case "Venus": {
      fuel = 300000
    }
    case "Mercury": {
      fuel = 500000
    }
    case "Mars": {
      fuel = 700000
    }
    default: {
      fuel = 0
    }
  }
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("You are on, " planet)
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  int fuelRemaining, fuelCost

  fuelRemaining = fuel

  fuelCost = calculateFuel(planet)
}


func main() {
  // Test your functions!
  
  // Create `planetChoice` and `fuel`
  
  // And then liftoff!
  
}