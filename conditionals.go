package main

import (
	"fmt"
	"time"
)

func main() {

	if 5%1 == 0 {
		fmt.Println("5 is divisible by 1")
	}

	number := -9988

	if number < 0 {
		fmt.Println(number, "is negative")
	} else if number < 10 {
		fmt.Println(number, "has 1 digit")
	} else {
		fmt.Println(number, "has multiple digits")
	}

	/*
		Switch case Conditionals
	*/

	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday, time.Friday:
        fmt.Println("It's the weekend",time.Now().Day())
    default:
        fmt.Println("It's a weekday", time.Now().Day())
    }

}
