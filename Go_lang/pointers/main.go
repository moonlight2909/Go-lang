package main

import "fmt"

func modifyvaluebyreference(a *int) {
	*a = *a * 20
}
func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := 2
	ptr := &num

	fmt.Println("Num has value :", num)
	fmt.Println("Ptr has value : ", ptr)
	fmt.Println("Data contains : ", *ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 10
	modifyvaluebyreference(&value)
	fmt.Println("Value contains :", value)

}
