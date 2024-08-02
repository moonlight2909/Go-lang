package main

import "fmt"

func simpleFunction() {
	fmt.Println("this is a simple function")
}
func add(a, b int) (result int) {
	result = a + b
	return result
}
func main() {
	fmt.Println("hii we are learning Go Lang")
	simpleFunction()
	ans := add(5, 7)
	fmt.Println("add of two number", ans)
}
