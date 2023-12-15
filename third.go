package main

import (
	"fmt"
	"math"
)

func saygreet(n string) {
	fmt.Printf("Good Morning @ %v \n", n)
}
func byegreet(n string) {
	fmt.Printf("Good Bye @ %v \n", n)
}
func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}
func cirarea(no float64) float64 {
	return math.Pi * no * no
}

func main() {
	saygreet("Alice")
	saygreet("Luther")
	byegreet("Alice and Luther")

	cycleName([]string{"harsh", "vardhan", "saxena"}, saygreet)
	cycleName([]string{"harsh", "vardhan", "saxena"}, byegreet)

	fmt.Println(cirarea(32.5))
}
