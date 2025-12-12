package main

import "fmt"

type Student struct {
	Name string
	Age int
	IsPass bool
}

func printStudentDegree(student *Student){
	if student.IsPass{
		student.Name = student.Name + " S.Kom"
	}
}

func (student *Student) PrintStudentDegree(){
	if student.IsPass{
		student.Name = student.Name + " S.Kom"
	}
}

func main(){
	// numberA := 25
	// numberB := &numberA

	// fmt.Println("Value of numberA:", numberA)
	// fmt.Println("Address of numberA:", numberB)
	// fmt.Println("Value of numberB:", *numberB)

	// var numberA int = 25
	// var numberB *int = &numberA
	// fmt.Println("Value of numberA:", numberA)
	// fmt.Println("Address of numberA:", numberB)
	// fmt.Println("Value of numberB:", *numberB)

	// numberA = 30
	// *numberB = numberA
	// fmt.Println("Value of numberA:", numberA)
	// fmt.Println("Address of numberA:", numberB)
	// fmt.Println("Value of numberB:", *numberB)

	// value := 50
	// fmt.Println("Old value:", value)
	// changeValue(&value, 255)
	// fmt.Println("New value:", value)

	student := Student{
		Name: "Budi",
		Age: 22,
		IsPass: true,
	}
	// printStudentDegree(&student)
	// fmt.Println("Student Name using function:", student.Name)

	student.PrintStudentDegree()
	fmt.Println("Student Name using method:", student.Name)
}

// func changeValue(oldValue *int, newValue int){
// 	*oldValue = newValue
// }
