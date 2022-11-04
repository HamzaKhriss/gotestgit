package main

import "fmt"

func main() {
	var a, b int
	for a = 0; a < 8; a++ {
		for b = a + 1; b != 10; b++ {
			fmt.Printf("%d%d, ", a, b)
		}
	}
	fmt.Print("89")
}
