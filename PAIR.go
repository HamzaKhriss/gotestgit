package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("enter a number : ")
	fmt.Scanln(&a)
	fmt.Print("enter a number : ")
	fmt.Scanln(&b)
	for i := 1; i < b-a; i++ {
		if (a+i)%2 == 0 {
			fmt.Println(a + i)
		}
	}

}
