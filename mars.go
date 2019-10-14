package main

import "fmt"

func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(65.0 * 0.3783)
	fmt.Print(" kg, and I would be ")
	fmt.Print(39 * 365 / 687)
	fmt.Println(" years old.")

	fmt.Println("My weight on the surface of Mars is", 65.0*0.3783,
		"kg, and I would be", 39*365/687, "years old.")

	fmt.Printf("My weight on the surface of Mars is %v kg, ", 65.0*0.3783)
	fmt.Printf("and I would be %v years old.\n", 39*365/687)
}
