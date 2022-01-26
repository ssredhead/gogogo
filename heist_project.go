package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
rand.Seed(time.Now().UnixNano())
eludedGuards := rand.Intn(100)

isHeistOn := true

if eludedGuards >= 50 {
  fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
} else {
  isHeistOn = false
  fmt.Println("Plan a better disguise next time?")
}

openedVault := rand.Intn(100)

if isHeistOn && openedVault >= 70 {
  fmt.Println("Grab and GO!")
} else if isHeistOn {
  isHeistOn = false
  fmt.Println("Vault cannot be opened.")
}

leftSafely := rand.Intn(3)

if isHeistOn {
  switch leftSafely {
    case 0: {
      isHeistOn = false
      fmt.Println("Heist failed")
    }
    case 1: {
      isHeistOn = false
      fmt.Println("Woops, caught at the door")
    }
    default: {
      fmt.Println("Start the getaway car!")
    }
  }
  if isHeistOn {
    amtStolen := 1000 + rand.Intn(1000000)
    fmt.Println("Amount stolen: $", amtStolen)
  }
}

fmt.Println("isHeistOn is currently:", isHeistOn)
}
