package main

import "fmt"

func main() {
	var a string
	var flag int
	fmt.Print("enter a word : ")
	fmt.Scanln(&a)
	for i := 0; i < len(a)/2; i++ {
		if string(a[i]) == string(a[len(a)-1-i]) {
			flag += 1
		}
	}
	if flag == len(a)/2 {
		fmt.Print("it s a palin")
	} else {
		fmt.Print("it s not a palin")
	}
}
