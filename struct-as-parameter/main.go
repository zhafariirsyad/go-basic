package main

import "fmt"

type Person struct{
	Id int
	Name string
	Age int
	Email string
	IsActive bool
}

func structAsParameter(person Person){
	fmt.Println(person)
}

func main(){
	person := Person{
		Id:1,
		Name:"John Doe",
		Age:25,
		Email: "john.doe@example.com",
		IsActive: true,
	}

	structAsParameter(person)
}