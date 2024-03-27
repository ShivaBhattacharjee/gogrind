package main

import "fmt"

func main() {
	var name [6]string
	name[0] = "John"
	name[1] = "Paul"
	name[2] = "George"
	name[3] = "Ringo"
	println("The length of array is ", len(name))
	fmt.Println(name[0], name[1], name[2], name[3])

	var normies = [2]string{"John", "Paul"}
	fmt.Println("The normies are", normies[0], normies[1])
}
