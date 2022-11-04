package main

import "fmt"

func main() {
	var a, b, p1, p2, p3, reduc, total float64
	fmt.Printf("enter the price of the item choosen : ")
	fmt.Scanln(&a)
	fmt.Printf("enter the number of the item choosen : ")
	fmt.Scanln(&b)
	fmt.Printf("enter the discount amount : ")
	fmt.Scanln(&reduc)
	p1 = a * b
	fmt.Println("TOTAL HORS TAXE AVANT REMISE : ", p1)
	p2 = p1 - (p1 * reduc / 100)
	fmt.Println("total after reduc is : ", p2)
	p3 = p2 * 0.2
	fmt.Println("tva = ", p3)
	total = p2 + p3
	fmt.Println("the total to pay is : ", total)
}
