// maps in go are like dictionaries in python or objects in javascript
// maps are key value pairs
// maps are unordered collection of key value pairs
// maps are reference types
// maps are used to look up values by keys
package main

import "fmt"

func main() {
	students := make(map[int]string)
	students[1] = "John"
	students[2] = "Doe"
	students[69] = "Man I hate Enginnering Maths"
	fmt.Println(students)
	fmt.Println("Student with roll no 69 is", students[69])

	// delete values from a map

	delete(students, 69)

	fmt.Println("Studets now after deleting of roll no 69", students)
}
