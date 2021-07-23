package main

import (
    "fmt"
    "math"
)

const stringConst string = "Constant String Testing.."

func main() {

	/*
		Constant variable is always constant, it can't be modified and we can't 
		change the value of any constant variable.
	*/

	// stringConst = "assign any value" not works

    fmt.Println(stringConst)

    const number  = 12540

    const result = 3e20 / number

    fmt.Println(math.Sin(number))
}