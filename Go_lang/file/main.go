package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Enter while creating file ,", err)

	}
	defer file.Close()
	content := "hello world"
	_, error := io.WriteString(file, content)
	if error != nil {
		fmt.Println("Enter while creating file ,", error)
		return ;
	}
	fmt.Println("Successfully created ")
}
