package main

import "fmt"

func main() {
	var a, total int
	total = 0
	fmt.Print("write a number :")
	fmt.Scanln(&a)
	for i := a; i > 0; i-- {
		total = total + i*i
	}
	fmt.Println(total)
}
