package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 6967
	intNum += 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	floatNum += 1
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var inNum32 int32 = 2
	var result float32 = floatNum32 / float32(inNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int =2
	fmt.Println(intNum1/intNum2) //will give only quotient upto decimal only the whole no part after decimal will be discarded
	fmt.Println(intNum1%intNum2) //will give only remainder

	var Mystring1 string = "acha"
	var Mystring2 string = "beta"
	fmt.Println(Mystring1+Mystring2)

	var mystring2 string = "hello"+"-"+"world"
	fmt.Println(mystring2)

	fmt.Println("length of mystring2 is", len(mystring2))
	fmt.Println("Rune count of mystring2 is", utf8.RuneCountInString(mystring2)) //difference btw len() and RuneCountInString() is that len() gives the number of bytes and RuneCountInString() gives the number of characters

	//here len() gives the number of bytes and RuneCountInString() gives the number of characters
	//eg. "hello" has 5 characters and 5 bytes
	// "world" has 5 characters and 5 bytes
	// "hello-world" has 11 characters and 11 bytes

	s := "你好"
	fmt.Println(len(s)) // 6 bytes

	c := '😊'
	fmt.Printf("%T", c) // int32
	fmt.Println("\nvalue of c is",c) // 128522

	var i = 10 // int
	fmt.Println(i)

	var ch rune = 'A'
    fmt.Println(ch)        // 65
    fmt.Printf("%c\n", ch)   // A

	var myBool bool = true
	fmt.Println(myBool) //true
	myBool = false
	fmt.Println(myBool) //false

	var intNum3 rune
	fmt.Println(intNum3)

	myVar := "text"
	myVar += "lod"
	fmt.Println(myVar)

	var1, var2 := 1,2
	fmt.Println(var1, var2)

	const pi float32 = 3.1415
	fmt.Println(pi)

}

/*
6968
1.23456799e+07
5.05
1
1
achabeta
hello-world
length of mystring2 is 11
Rune count of mystring2 is 11
6
int32
value of c is 128522
10
65
A
true
false
0
textlod
1 2
3.1415
*/
