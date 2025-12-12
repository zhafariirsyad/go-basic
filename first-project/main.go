package main

import (
	"first-project/calculation"
	"fmt"
)

func main() {
	fmt.Println("Hello, Pai!")

	fmt.Println("This is result of Add Function")
	result := calculation.Add(10,69)
	fmt.Println(result)

	fmt.Println("This is result of Multiply Function")
	resultMultiply := calculation.Multiply(10,69)
	fmt.Println(resultMultiply)
}