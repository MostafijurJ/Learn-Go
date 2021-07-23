package main

import (
	"fmt"
)

//Single instance and single return types
func add(numberOne int, numberTwo int) int{
	var result = numberOne+numberTwo
	return result
}

//Multiple return type
func multipleReturn(one, two int)(add, sub, multi, div int){
	 add = one + two
	 multi = one * two
	 sub = one -two
	 div = one / two
	 return
}

/*
	N.B: If you want to access any functions or variable outside of the package 
	it names should start with capital latter, upper case or camel case. 

	roll: 9101083
	reg: 15226126324
*/

func main(){

	var sout  = add(5,6)
	fmt.Println("Single return types result : ", sout)

	var add,sub, multi, div = multipleReturn(1000,100)

	fmt.Println("catches multiple instances")
	fmt.Println("addition : ", add,"\t subtraction : ", sub, "\tmultiplication :", multi, "\tDivission : ", div)


}