package main

import "fmt"

func main(){
	var languages = [...]string{"Go", "Python", "JavaScript", "Java", "C++"}

	languages[4] = "TypeScript"

	for index, language := range languages {
		fmt.Println(index, language)
	}
}
