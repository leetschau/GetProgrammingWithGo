package main
import (
  "fmt"
)

func main() {
  const days = 38
  const hours = days * 24
  const distance = 56000000
  const least_speed = distance / hours
  fmt.Printf("You have to travel at least %vkm/h to reach Mars in 28 days.\n", least_speed)
}
