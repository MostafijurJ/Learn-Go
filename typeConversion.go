package main

import "fmt"
import "strconv"
//Import strconv package for String coversion

func main() {

	var integerNumber int = 56
	fmt.Printf("%v \t %T \n", integerNumber, integerNumber)

    var stringNumber  string
	//this method is call to convert int to string
	stringNumber  = strconv.Itoa(integerNumber)

	fmt.Printf("%v \t %T \n", stringNumber, stringNumber)

}
