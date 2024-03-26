package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time right now is ")
	currTime := time.Now()
	// the time formatting is done using the reference date Mon Jan 2 15:04:05 MST 2006
	fmt.Print(currTime.Format("01-02-2006 Monday 15:04:05"))

	// well how to add custom date and time

	cusDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Custom date is \n", cusDate.Format("01-02-2006 Monday 15:04:05"))
}
