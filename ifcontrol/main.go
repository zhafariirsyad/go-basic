package main

import "fmt"

func main(){
	age := 30
	var status string

	if age < 18 {
		status = "Minor"
	} else if age >= 18 && age < 65 {
		status = "Adult"
	} else {
		status = "Senior"
	}

	fmt.Println("Age:", age)
	fmt.Println("Status:", status)
}
