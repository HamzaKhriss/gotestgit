package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c, root1, root2, discriminant float64

	fmt.Print("Enter the a, b, c of Quadratic equation = ")
	fmt.Scanln(&a, &b, &c)

	discriminant = (b * b) - (4 * a * c)

	if a != 0 && discriminant > 0 {
		root1 = (-b + math.Sqrt(discriminant)/(2*a))
		root2 = (-b - math.Sqrt(discriminant)/(2*a))
		fmt.Println("Two Distinct Real Roots Exist: root1 = ", root1, " and root2 = ", root2)
	} else if a != 0 && discriminant == 0 {
		root1 = -b / (2 * a)
		fmt.Println("one unique solution exists root = ", root1)
	} else if a != 0 && discriminant < 0 {
		println("no solution")
	} else {
		println("a is equal to zero try again")
	}
}
