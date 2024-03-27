package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your age ")
	reader := bufio.NewReader(os.Stdin)
	number, _ := reader.ReadString('\n')
	age, err := strconv.ParseInt(strings.TrimSpace(number), 10, 64)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	if age >= 18 {
		println("You are an adult")
	} else {
		println("You are a kid")
	}
}
