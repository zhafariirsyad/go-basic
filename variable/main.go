package main

import "fmt"

func main(){
	// First way to declare variable
	// var name string
	// name = "Pai"
	var name string = "Pai"
	fmt.Printf("My name is : %s", name)

	fmt.Println("")

	// Second way to declare variable
	// := just for new variable declaration
	// = for assigning value to existing variable
	age := 24
	fmt.Printf("My age is : %d", age)
	age = 25
	fmt.Printf("\nMy new age is : %d", age)
}