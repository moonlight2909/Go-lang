package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}
type Contact struct {
	Email string
	Phone string
}
type Address struct {
	House int
	Area  string
	State string
}
type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
	// EmployeeID         int
	// EmployeeSalary     float64
	// EmployeeDepartment string
	// EmployeeNumber     string

}

func main() {
	var himanshu Person
	//fmt.Println("Himanshu Person :", himanshu)
	himanshu.Firstname = "Himanshu"
	himanshu.Lastname = "Kashyap"
	himanshu.Age = 19

	//fmt.Println("Details of Himanshu : ", himanshu)
	// 2 method
	person1 := Person{
		Firstname: "Aakash",
		Lastname:  "kashyap",
		Age:       19,
	}
	fmt.Println("person1 details :", person1)

	// using new keyword
	var person2 = new(Person)
	person2.Firstname = "himanshu"
	person2.Lastname = "kashyap"

	var Employee1 Employee
	Employee1.Person_Details = Person{
		Firstname: "Himanshu",
		Lastname:  "kashyap",
		Age:       19,
	}
	Employee1.Person_Contact = Contact{
		Email: "himanshukashyap98054@gmail.com",
		Phone: "8544721993",
	}
	Employee1.Person_Address = Address{
		House: 1,
		Area:  "kangra",
		State: "Himachal",
	}

	fmt.Println("Employee 1 : ", Employee1)
}
