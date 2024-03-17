package main

import "fmt"

type User struct {
	name    string
	surname string
}

func main() {
	// Maps work is a key-value pair data structure, similar to an object in javascript or map in solidity.
	users := make(map[string]User)

	fmt.Println(users)

	users["Karim"] = User{"Karim", "Yassine"}
	users["Madalina"] = User{"Madalina", "Niculescu"}
	users["Nadia"] = User{"Nadia", "Yassine"}

	fmt.Println(users)
	fmt.Println(users["Karim"].surname)

	// Deletes specific key
	delete(users, "Madalina")
	fmt.Println(users)

	m, ok := users["Madalina"]
	fmt.Printf("Value[%d] Ok[%d] \n", m, ok)

	// When looping over a map, the order is random.
	// This is done propositaly by go to not allotw engineers to create a for loop based on the order of a map.
	for u := range users {
		fmt.Println(u)
	}

	// write about value change

	// This wont work, you need to change the value and re-assign it.
	// *users["Karim"].name = "NOT"

	// This will work.
	// 1. check if ok is true (key exists in map)
	// 2. u is a copy. change the value of the copy.
	// 3. assign the copy to the key inside map.
	if u, ok := users["Nadia"]; ok {
		u.surname = "Niculescu"
		users["Nadia"] = u
	}

	fmt.Println(users)
}
