package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Number is :", i)
	}

	// infinite loop using break
	counter := 0
	for {
		fmt.Println("Infinite loop")
		counter++
		if counter == 3 {
			break
		}

	}

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, value := range number {
		fmt.Printf("Index of number is %d , and values is %d\n", index, value)
	}

	data := "himanshu kashyap"
	for index, char := range data {
		fmt.Printf("Index of data is %d , and values is %c\n", index, char)
	}
}
