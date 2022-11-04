package main

import "fmt"

func main() {
	var a, total int
	total = 1
	fmt.Print("write a number :")
	fmt.Scanln(&a)
	for i := a; i > 0; i-- {
		total = total * i
	}
	fmt.Println(total)
}
