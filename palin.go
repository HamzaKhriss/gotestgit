package main

import "fmt"

func main() {
	var a, b string
	var flag int
	fmt.Print("enter a word : ")
	fmt.Scanln(&a)
	fmt.Print("enter a second	 word : ")
	fmt.Scanln(&b)
	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			flag += 1
		}
	}
	fmt.Println("the hamming difference is : ", flag)
}
