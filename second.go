package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello Google")
	var name1 string = "Harsh"
	fmt.Println(name1)
	name2 := "damncat"

	fmt.Println(name2)

	var num1 int = 9
	num2 := 10
	fmt.Println(num1, num2)

	var str = fmt.Sprintf("the texted version of two numers is represented as %v", num1+num2)
	fmt.Println(str)

	fmt.Println(strings.Contains("harshvardhan", "harsh"))

}
