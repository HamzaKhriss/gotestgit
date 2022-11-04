package main

import "fmt"

func main() {
	for b := 0; b < 10; b++ {
		for i := 0; i < 9; i++ {
			fmt.Printf("%d%d, ", b, i)
		}
		fmt.Printf("%d9\n", b)
	}

}
