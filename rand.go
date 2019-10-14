package main

import (
  "fmt"
  "math/rand"
)

func main() {
  var num = rand.Intn(10) + 1
  fmt.Println(num)
  num = rand.Intn(10) + 1  // 不用写成 math/rand
  fmt.Println(num)
}
