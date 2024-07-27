package main

import "fmt"

// func divide1(a, b float64) float64 {
// 	return a / b
// }
// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator must not be Zero")
// 	}
// 	return a / b, nil
// }
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominator must not be Zero"
	}
	return a / b, "nil"
}
func main() {
	fmt.Println("this is code of error handling")
	ans, _ := divide(10, 0)
	fmt.Println("your ans is ", ans)

}
