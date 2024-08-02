package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("starting of the project")
	data := add(5, 6)
	defer fmt.Println("Data is : ", data)
	defer fmt.Println("middle of the project")
	fmt.Println("end of the project")

}
