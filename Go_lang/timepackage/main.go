package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time :", currentTime)

	fmt.Printf("Type of current time is %T\n", currentTime)

	// formated := currentTime.Format("dd-mm-yyyy")
	formated := currentTime.Format("02-01-2006, Monday ,  3:04 PM")

	fmt.Println("Formatted time  :", formated)

	// layoutstring = "2006-01-02"
	// datastr := "2033-11-25"
	//formateed := time.Parse(layoutstring, datastr)

	//fmt.Println(formateed)

}
