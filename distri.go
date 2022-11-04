package main

import "fmt"

func main() {
	var B200, B100, B50, B20, P10, P5, P2, P1, S int
	fmt.Printf("Enter number how much : ")
	fmt.Scanf("%d", &S)
	B200 = S / 200
	B100 = (S % 200) / 100
	B50 = (S % 200 % 100) / 50
	B20 = (S % 200 % 100 % 50) / 20
	P10 = (S % 200 % 100 % 50 % 20) / 10
	P5 = (S % 200 % 100 % 50 % 20 % 10) / 5
	P2 = (S % 200 % 100 % 50 % 20 % 10 % 5) / 2
	P1 = (S % 200 % 100 % 50 % 20 % 10 % 5 % 1) / 1
	fmt.Printf("%d billet de 200, %d billet de 100, %d billet de 50, %d billet de 20, %d piece de 10, %d piece de 5, %d piece de 2, %d piece de 1\n", B200, B100, B50, B20, P10, P5, P2, P1)
	bruh()
}

func bruh() {
	fmt.Println("bruhhhhhhhhhhhhhhhhhhh")
}
