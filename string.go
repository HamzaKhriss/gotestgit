package main

import (
	"fmt"
)

func main() {
	var b string
	fmt.Print("enter a word : ")
	fmt.Scanln(&b)
	for i := 0; i < len(b); i++ {
		fmt.Println(string(b[i]))
	}
}
