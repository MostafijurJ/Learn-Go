package main

import "fmt"

func main() {

	fmt.Println("\nPrinting Booleans \n")
	//Boolean
	//By default boolean initialization is false
	var booldef  bool
	fmt.Printf("Deafult value: %v \t %T\n", booldef, booldef)

	var boolean bool  = false
	fmt.Printf("Initialized value : %v \t %T\n", boolean, boolean)

	//Initialization and conditional testing
	boolTest := 1 ==1
	fmt.Printf("Conditional testing : %v\n",boolTest)

	// Integer
	/* 
		these are mainly 4 types
		int8 range  -128  to 127
		int16 range -32768 to 32767
		int32 range -2147483648 to 2147483647
		int 64 range maximum 

		N.B for any types of operation you should need to same type of operands or you have to cast it as same type
	*/

	fmt.Println("\nPrinting Integers \n")
	var in8  int8 = 12
	fmt.Printf("int8 %v \t %T \n", in8, in8)
	 var a int  = 10
	 var b int8 = 5
	 /*
		it won't works cause of type mismatch.
		sum := a+b 

		so you have to cast it. 
	 */
	 sum := a+int(b)

	 fmt.Printf("types issue : %v\n", sum)






}
