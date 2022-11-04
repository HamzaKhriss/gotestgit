package main

import "fmt"

func main() {
	var n int
	fmt.Printf("Enter number of rows: ")
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		for s := i; s < n; s++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= (2*i - 1); j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
