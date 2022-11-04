package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("enter a :")
	fmt.Scanln(&a)
	fmt.Print("enter b :")
	fmt.Scanln(&b)
	for i := 0; i < b; i++ {
		c += a
	}
	fmt.Println("the result is ", c)
}
