package main

import "fmt"

func main() {
	var intArr [3]int32 //empty arr of len 3
	intArr[0] = 123     //adding values to the array at 0th index
	fmt.Println(intArr)
	fmt.Println(intArr[1:3]) //slicing the array from index 1 to 3

	fmt.Println(&intArr[0]) //print memory address
	fmt.Println(&intArr[1]) //print memory address
	fmt.Println(&intArr[2]) //print memory address

	intArr2 := [...]int32{1, 2, 3} // auto len of array will be 3
	fmt.Println(intArr2)

	var intSlice []int32 = []int32{4, 5, 6} //preinitiallise a slice with len n cap 3 both
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) //incr cap to 6
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9} //preinitiallise a slice with len n cap 2 both
	intSlice = append(intSlice, intSlice2...)
	fmt.Println("\n", intSlice)

	var intSlice3 []int32 = make([]int32, 3, 10) //preinitiallise a slice with len 3 and cap 10
	fmt.Println(intSlice3)                       // [0 0 0]

	var myMap map[string]uint8 = make(map[string]uint8) //prein empty map
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"adam": 23, "sarah": 45} //preinit map
	fmt.Println(myMap2["adam"])
	var age, ok = myMap2["adam"] //ok check if the key is present or not
	delete(myMap2, "adam")       // delete key value from map
	if ok == true {
		fmt.Printf("umur adam %v \n", age)
	} else {
		fmt.Println("ga ada adam \n")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, age: %v \n", name, age)
	} //iterating map 2 by key value pairs

	for i, v := range intArr {
		fmt.Printf("Index: %v, value: %v \n", i, v)
	} //iterating array by index value pairs
}

/*
[123 0 0]
[0 0]
0x291cab2a80e0
0x291cab2a80e4
0x291cab2a80e8
[1 2 3]
The length is 3 with capacity 3The length is 4 with capacity 6
 [4 5 6 7 8 9]
[0 0 0]
map[]
23
umur adam 23
Name: sarah, age: 45
Index: 0, value: 123
Index: 1, value: 0
Index: 2, value: 0
*/
