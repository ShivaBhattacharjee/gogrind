// store user input in a variable
// we use bufio package to read user input
// bufio package is used to read user input from the console
// os package is used to read user input from the console

// This is called comma error syntax
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Printf("Type of %T  \n", name)
	reader := bufio.NewReader(os.Stdin)
	print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello", name)
	fmt.Printf("Type of reader is %T \n", reader)
}
