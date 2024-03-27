package main

import (
	"fmt"
	"sort"
)

func main() {
	var students = []string{"John", "Paul", "George", "Ringo"}
	// in order to add values in a slice we use append method
	students = append(students, "Pete", "Stu")
	fmt.Printf("The type of students is %T\n", students)
	fmt.Println("The students are", students)

	fmt.Println("well some append methods with indexing")

	students = append(students[1:3]) //similary we cann mess around with arguments to get values from a slice
	fmt.Println("The students are", students)

	// some sorting related stuff

	marks := make([]int, 4)
	marks[0] = 12
	marks[1] = 34
	marks[2] = 56
	marks[3] = 78
	// marks[4] = 12 // This is an error as the length of the slice is 4
	// however we can use append method to add values to the slice
	marks = append(marks, 69)
	fmt.Println("The marks are", marks)
	isSorted1st := sort.IntsAreSorted(marks)
	fmt.Println("Is the array sorted?", isSorted1st)
	sort.Ints(marks)
	fmt.Println("The sorted marks are", marks)
	// Check if array is sorted or not
	isSorted := sort.IntsAreSorted(marks)
	fmt.Println("Is the array sorted?", isSorted)

	// remove the values from the slice based upon index

	var languages = []string{"C", "C++", "Java", "Python", "Go", "Ruby", "JavaScript"}
	fmt.Println("The languages are", languages)
	// append can add as well as remove values from the slice

	const indexToRemove int = 3

	languages = append(languages[:indexToRemove], languages[indexToRemove+1:]...) //This removes thte 3rd index from the slice

	fmt.Println("The languages are after removal ", languages)
}
