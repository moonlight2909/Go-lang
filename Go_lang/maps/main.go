package main

import "fmt"

func main() {
	studentsgrade := make(map[string]int)
	studentsgrade["himanshu"] = 90
	studentsgrade["raj"] = 80
	studentsgrade["Anubhav"] = 88

	fmt.Println("Marks of himanshu :", studentsgrade["himanshu"])
	studentsgrade["himanshu"] = 100
	fmt.Println(" New Marks of himanshu :", studentsgrade["himanshu"])

	delete(studentsgrade, "himanshu")
	//fmt.Println("Marks of himanshu :", studentsgrade["himanshu"])
	// check if key is present or not
	grades, exists := studentsgrade["David"]
	fmt.Println("Marks of David :", exists)
	fmt.Println("Marks of David :", grades)

	for index, value := range studentsgrade {
		fmt.Printf("key is %s and marks is %d/n", index, value)
	}

	person := map[string]int{
		"Alice ": 1100,
		"Bob ":   200,
		"chi ":   300,
	}
	for index, value := range person {
		fmt.Printf("key is %s and marks is %d/n", index, value)
	}
}
