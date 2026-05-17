package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé") //if not rune then it will skip 2 index as it sees é as 2 characters and 1 byte each
	var indexed = myString[1]
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString{
		fmt.Println(i, v) //gives value in terms of ascii and index in terms of bytes
	}
	fmt.Printf("\n the length of string is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\n myRune = %v", myRune)

	var strSlice = []string{"s","u","b","s","c","r","i","b","e"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i]) //append string in string builder
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}