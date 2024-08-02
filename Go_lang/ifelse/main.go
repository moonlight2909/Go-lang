package main

import "fmt"

func check(a, b int) {
	if a < b {
		fmt.Println("a is smaller then b")
	} else {
		fmt.Println("b is smaller then a ")
	}
}
func main() {
	var x int
	fmt.Scan(&x)
	if x > 5 {
		fmt.Println("x is greater then 5")
	} else {
		fmt.Println("x is smaller then 5")
	}
	check(3, 8)

}
