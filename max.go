package main

import (
	"fmt"
	"math"
)

func main() {
	var x, a, max, maxposition float64
	fmt.Println("please enter a number : ")
	fmt.Scanln(&a)
	var j float64 = 1
	for {
		fmt.Scan(&x)
		max = math.Max(a, x)
		j += 1
		if max == x {
			a = x
			maxposition = j
		}
	}
	fmt.Println("the maximum is : ", a, "position", maxposition)
}
