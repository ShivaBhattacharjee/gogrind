// well what i can conclude is converting string to a number is a pain in the ass atleas in go lang
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conversion Enter a number")

	reader := bufio.NewReader(os.Stdin)

	// ReadString will read until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter.
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
	fmt.Printf("The type of input is: %T \n ", input)

	// in order to convert the input to a number we need to use the strconv package

	conver, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		conver += 20
		fmt.Println("The converted number is ", conver)
	}
}
