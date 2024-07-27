package main

import "fmt"

// Slice act like a vector
func main() {
	// number := []int{1, 2, 3, 4, 5}
	// fmt.Println("Number :", number)

	// fmt.Println("length if this :", len(number))
	// // format
	// name := []string{}
	// fmt.Println("name is", name)
	// //name := append(name, 1, 2, 4, 5, 6, 7)

	number := make([]int, 3, 5)

	number = append(number, 4)
	number = append(number, 5)
	number = append(number, 6)
	fmt.Println("Slice :", number)
	fmt.Println("length :", len(number))
	fmt.Println("capacity :", cap(number))

}
