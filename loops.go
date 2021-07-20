package main

import "fmt"

func main(){

	fmt.Print("Looping\n")
	/*
		IN Go all the looping operation can be done by only for loop
	*/
	
	for i := 0; i < 5; i++ {
		fmt.Println("Loop ", i+1)
	}

}
