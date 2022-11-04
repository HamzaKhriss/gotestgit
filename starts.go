package main

import (
	"fmt"
)

func main() {
	var b int
	fmt.Print("please enter a number : ")
	fmt.Scanln(&b)
	for i := b; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
