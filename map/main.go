package main

import "fmt"

func main(){
	// fmt.Println("Hello, World!")

	// var myMap = map[string]string{
	// 	"name":    "John Doe",
	// 	"address": "123 Main St",
	// 	"email":   "john.doe@example.com",
	// }

	// for key, value := range myMap {
	// 	fmt.Println(key, ":", value)
	// }

	// test, isAvailable := myMap["nam"]
	// fmt.Println("Name:", test)
	// fmt.Println("Is Available:", isAvailable)


	// var secondMap map[string]int

	// secondMap = map[string]int{}

	// secondMap["Math"] = 10

	// fmt.Println(secondMap["Math"])


	// map slice
	var mapSlice = []map[string]string{
		{
			"name": "Pai",
			"age":  "21",
		},
		{
			"name": "Piwey",
			"age":  "20",
		},
	}

	fmt.Println(mapSlice)
}