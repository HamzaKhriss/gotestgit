package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	a := bufio.NewReader(os.Stdin)
	var c string
	fmt.Print("enter a word : ")
	line, _ := a.ReadString('\n')
	b := strings.ReplaceAll(line, " ", "")
	for i := len(b) - 1; i >= 0; i-- {
		c = c + string(b[i])

	}

	fmt.Println(reflect.TypeOf(c), len(c))
	fmt.Println(reflect.TypeOf(b), len(b))

	if c == b {
		fmt.Print("it s a palin")

	} else {

		fmt.Print("sdhgf")

	}

}
