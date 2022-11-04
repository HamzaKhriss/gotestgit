package main

import "fmt"

func main() {
	var j, m, a int
	fmt.Print("enter the day : ")
	fmt.Scanln(&j)
	fmt.Print("enter the month : ")
	fmt.Scanln(&m)
	fmt.Print("enter the year : ")
	fmt.Scanln(&a)

	switch true {
	case a%4 == 0 && m == 2 && j == 28:
		fmt.Println("the date is : 29/02/", a)
	case a%4 != 0 && m == 2 && j == 28:
		fmt.Println("the date is : 01/03/", a)
	case (m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12) && j < 31:
		fmt.Println("the date is ", j+1, "/", m, "/", a)
	case (m == 4 || m == 6 || m == 9 || m == 11) && j < 30:
		fmt.Println("the date is ", j+1, "/", m, "/", a)
	case j == 31 && m == 12:
		fmt.Println("the date is 01/01/", a+1)
	default:
		fmt.Println("the date is 01/", m+1, "/", a)
	}
}
