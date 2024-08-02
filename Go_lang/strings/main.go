package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two two one three ont two"
	count := strings.Count(str, "two")
	fmt.Println("Count of two : ", count)

	str1 := "Hello World"
	trimed := strings.TrimSpace(str1)
	fmt.Println("trimmed : ", trimed)

}
