package main

import "fmt"

// Public variables should be declared with a capital letter
const ThisIsPublic = "Shiva"

func main() {

	/*
		The primitive data stucture in go consists of:
		1. String
		2. Integers
		 Integers are further divided into:
		 a. Signed Integers :int(platform dependent) ,  int8, int16, int32, int64
		 b. Unsigned Integers :unint(platform dependent if system is 32 bit it will assign 32 bit if higher it will update accordingly) ,
		  uint8, uint16, uint32, uint64

		3. Boolean (true or false)
		4. Float
			Float are further divided into:
			a. Float32
			b. Float64
	*/

	/*
		Non primitive data types in go are:
		*diff between array and slice is that array has fixed length and slice has dynamic length*
		1. Array , Array is a collection of similar data types
		2. Slice , Slice is a collection of similar data types
		3. Map , Map is a collection of key value pairs
		4. Struct , Struct is a collection of different data types
		5. Pointer , Pointer is a variable that stores the memory address of another variable
	*/

	// variables can be declared by using var and const keyword along with the data type
	var name string = "Shiva"
	fmt.Println("Hello world from ", name)

	// implementing the integer data type
	// usign int
	var age int = 25
	fmt.Println("Age is ", age)
	// using signed and unsigned integers
	// signed integers
	var smallNum int8 = 25 // -128 to 127 *range of signed int is -2^n to 2^(n-1)-1 *
	fmt.Println("Small number is ", smallNum)
	var bigNum int64 = 25 // -9223372036854775808 to 9223372036854775807
	fmt.Println("Big number is ", bigNum)

	// unsigned integers
	var smallNum1 uint8 = 25 // 0 to 255 *range of unsigned int is 0 to 2^n-1 *
	fmt.Println("Small number is ", smallNum1)
	var bigNum1 uint64 = 25 // 0 to 18446744073709551615
	fmt.Println("Big number is ", bigNum1)

	// empty variables
	var emptyString string
	fmt.Println("Empty string is ", emptyString)

	// implict type declaration
	var myWebsite = "https://www.theshiva.xyz" //lexers auto adds the type of the variable
	fmt.Println("My website is ", myWebsite)

	// Declaring variables without var
	// := is used to declare and initialize the variable
	test := "wow this operator is cool"
	fmt.Println(test)

	// %T is used to print the type of the variable
	fmt.Printf("Type of public variable is: %T", ThisIsPublic)
}
