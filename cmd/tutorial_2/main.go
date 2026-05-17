// functions, if else elif, switch case

package main

import(
	"fmt"
	"errors"
)


func printme(mystring string){
	fmt.Println("hello" + mystring)
}

func intDivision(num int, deno int) (int, int, error){
	var err error
	if deno == 0{
		err = errors.New("can't divide by zero")
		return 0, 0, err
	}

	var result int = num/deno
	var remainder = num%deno
	return result, remainder, err

}

func main(){
	ggstring := "nigh"
	printme(ggstring)
	var num=11
	var deno =2
	var result, remainder, err = intDivision(num, deno)
	if err!=nil{
		fmt.Printf(err.Error())
	}else if remainder==0{
		fmt.Printf("the result of the int is %v", result)
	}else{
		fmt.Printf("the result of the int is %v and the remainder is %v", result, remainder)
	}

	switch remainder{
	case 0:
		fmt.Println("\nThe Division was exact")
	case 1,2:
		fmt.Println("\nThe Division was close")
	default:
		fmt.Println("\nThe Division was not close")
	}
}

