package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const result = 59
	var guess = rand.Intn(100)
	for {
    fmt.Printf("You guess: %v\n", guess)
		if result == guess {
			fmt.Printf("You get it: %v\n", result)
      break
		} else if guess > result {
      fmt.Println("More!")
			guess = rand.Intn(guess)
		} else {
      fmt.Println("Less!")
			guess = guess + rand.Intn(100 - guess)
		}
	}
}
