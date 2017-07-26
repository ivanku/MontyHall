package main

import "fmt"
import "math/rand"
import "time"

const NoOfGames = 10000
const NoOfDoors = 3

func main() {
  fmt.Printf("Monty Hall %d-door experiment\n", NoOfDoors)

  s := rand.NewSource(time.Now().Unix())
  r := rand.New(s)

  switchStrategyWins := 0
  holdStrategyWins := 0

  for i := 0; i < NoOfGames; i++ {
    // randomly choose a door behind which there's a prize
    prizeDoor := r.Intn(NoOfDoors)

    // player picks a door at random
    playerChoice := r.Intn(NoOfDoors)
    fmt.Printf("Prize door: %d, player choice: %d", prizeDoor, playerChoice)

    // host randomly opens one door (which is not the player chosen door and not the prize door)
    var openDoor int
    for openDoor = r.Intn(NoOfDoors); openDoor == playerChoice || openDoor == prizeDoor; openDoor = r.Intn(NoOfDoors) {}
    fmt.Printf(", host opens door: %d", openDoor)

    if (playerChoice == prizeDoor) { 
      // first strategy - don't change the original choice
      holdStrategyWins++
      fmt.Printf(" - HOLD wins\n")
    } else {
      // second strategy - change the choice to another door
      var newChoice int
      for newChoice = r.Intn(NoOfDoors); newChoice == playerChoice || newChoice == openDoor; newChoice = r.Intn(NoOfDoors) {}
      if (newChoice == prizeDoor) {
        switchStrategyWins++
        fmt.Printf(" - SWITCH wins\n")
      } else {
        fmt.Printf(" - NO STRATEGY wins\n")
      }
    }
  }

  fmt.Printf("Hold strategy wins: %d, win probability: %3.2f\n", holdStrategyWins, float32(holdStrategyWins) / NoOfGames)
  fmt.Printf("Switch strategy wins: %d, win probability: %3.2f\n", switchStrategyWins, float32(switchStrategyWins) / NoOfGames)
}
