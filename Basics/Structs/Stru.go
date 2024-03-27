package main

import "fmt"

func main() {
	type Student struct {
		name   string
		rollNo uint64
		grade  string
	}

	sudent := Student{"John", 1, "A"}
	fmt.Printf("The details are %+v\n", sudent)
}
