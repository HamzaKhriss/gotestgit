package main

import "fmt"

func main() {
	var a int
	fmt.Print("write a number :")
	fmt.Scanln(&a)
	for i := 1; i < 11; i++ {
		fmt.Println(a*a + i)
	}

}
