package main

import "fmt"

func main(){
	// for
	// for i := 1; i <= 2; i++ {
	// 	fmt.Println("Iteration:", i)
	// }

	// while
	// i := 1
	// for i <= 2 {
	// 	fmt.Println("While Iteration:", i)
	// 	i++
	// }

	letter := "Panggilan saya adalah Pai"
	// for index, char := range letter {
	// 	fmt.Println("Index: ", index, "Character: ", string(char))
	// }

	// If index is not needed (use _)
	// for _, char := range letter {
	// 	fmt.Println("Character: ", string(char))
	// }

	for index, char := range letter {
		if index % 2 == 0{
			fmt.Println("Index: ", index, "Character: ", string(char))
		}
	}

	for _, char := range letter {
		letterString := string(char)

		switch letterString{
			case "a", "i", "u", "e", "o", "A", "I", "U", "E", "O":
				fmt.Println("Vowel:", letterString)
			default:
				fmt.Println("Consonant:", letterString)
		}
	}
}
