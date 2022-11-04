package main

import "fmt"

func main() {
	var b, c string
	fmt.Print("enter a word : ")
	fmt.Scanln(&b)
	for i := len(b) - 1; i >= 0; i-- {
		c = c + string(b[i])

	}
	fmt.Print(c)
}
