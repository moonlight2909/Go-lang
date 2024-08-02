package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 45
	fmt.Println("Number is : ", num)
	fmt.Printf("type of num is %T\n :", num)

	var data float64 = float64(num)
	data = data + 1.23
	fmt.Printf("Type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is : ", str)
	fmt.Printf("Type of str is %T\n", str)

	// number_string = "12345";
	// number_int := strconv.Atoi(number_string);

}
