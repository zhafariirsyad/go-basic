package main

import "fmt"

type Person struct {
	Id int
	Name string
	Age int
	Email string
	IsActive bool
}
// Embedding Person struct within Group struct
type Group struct {
	Name string
	Admin Person
	Members []Person
	IsAvailable bool
}

// This function called method as it is associated with Group struct
func (group Group) DisplayGroup(){
	fmt.Printf("Group Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Admin: %s", group.Admin.Name)
	fmt.Println("")
	fmt.Printf("Members: ")
	for _, member := range group.Members {
		fmt.Printf("%s, ", member.Name)
	}
}

func main(){
	person := Person{
		Id:1,
		Name:"John Doe",
		Age:25,
		Email: "john.doe@example.com",
		IsActive: true,
	}
	person2 := Person{
		Id:2,
		Name:"Jane Smith",
		Age:30,
		Email: "jane.smith@example.com",
		IsActive: false,
	}

	group := Group{
		Name : "Developers",
		Admin: person,
		Members: []Person{person,person2},
		IsAvailable: true,
	}

	group.DisplayGroup()
}

func displayGroup(group Group){
	fmt.Println(group)
	fmt.Printf("Group Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Admin: %s", group.Admin.Name)
	fmt.Println("")
	fmt.Printf("Members: ")
	for _, member := range group.Members {
		fmt.Printf("%s, ", member.Name)
	}
}
