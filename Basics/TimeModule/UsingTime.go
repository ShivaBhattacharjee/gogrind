package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time right now is ")
	currTime := time.Now()
	// the exact date and time need to be specified for the time formatting
	fmt.Print(currTime.Format("01-02-2006 Monday 15:04:05"))
}
