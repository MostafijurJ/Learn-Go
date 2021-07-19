package main

import (
	"fmt"
)

func main() {

	fmt.Println("\nPrinting Booleans \n")
	//Boolean
	//By default boolean initialization is false
	var booldef  bool
	fmt.Printf("Deafult value: %v \t  Type: %T\n", booldef, booldef)

	var boolean bool  = false
	fmt.Printf("Initialized value : %v \t  Type: %T\n", boolean, boolean)

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
	fmt.Printf("int8 %v \t  Type: %T \n", in8, in8)
	 var a int  = 10
	 var b int8 = 5
	 /*
		it won't works cause of type mismatch.
		sum := a+b 

		so you have to cast it. 
	 */
	 sum := a+int(b)

	 fmt.Printf("types issue : %v\n", sum)



	// Floating
	/* 
		these are mainly 2 types

		float32 range +-1.18E-38 to +-3.4E38
		float64 range +-2.23E-208 to +-1.8E308

		N.B: there are no bitwise operator for floating point number.  Bitwise operator is only work on integer value. 

	*/

	fmt.Println("\nPrinting Floating \n")

	var f32 float32 = 3.23
	 f32 = 1.23e5
	
	fmt.Println(f32)



	// Complex
	/* 
		Complex number is mainly two types
		complec128, complex64

		IN complex by number real operation is used for finding the real part of this nunmber

		N.B: Modulus operation is not available in complex number system
	*/

	fmt.Println("\nPrinting Complex \n")

	var cxa  complex64 = 1+9i
	var cxb complex64 = 2+7i
	fmt.Printf("Summation of Complex Number is %v \n", cxa+cxb)
	fmt.Printf("Abstruction of Complex Number is %v \n", cxa-cxb)
	fmt.Printf("Multiplication of Complex Number is %v \n", cxa*cxb)
	fmt.Printf("Division of Complex number's real part is %v \n", real(cxa/cxb))
	fmt.Printf("Division of Complex number's complex part is %v \n", complex(5,6))

	// Text 
	/* 

		IN complex by number real operation is used for finding the real part of this nunmber

		N.B: Modulus operation is not available in complex number system
	*/

	fmt.Println("\nText Types \n")
	

	 text := "Hello World"
	 //if we don't cast it, it will printed the ascii value of respected character.
	fmt.Printf("Text %v \t Tyep: %T \n", string(text[2]), text)

	// Convert a string into a byte slice or a set of bytes
	byteArr := [] byte(text)
	fmt.Printf("Byte Array %v \t Tyep: %T \n", byteArr, byteArr)

	/*
		rune is an alias for int32 and is equivalent to int32 in all ways
	*/

	runeTest := 'r'
	fmt.Printf("Rune  %v \t Tyep: %T \n", runeTest, runeTest)



}
