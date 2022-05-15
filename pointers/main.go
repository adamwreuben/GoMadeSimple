package main

import "fmt"

func main() {
	a := 5  //int
	b := &a // *int

	fmt.Print(a, b)

	//use * to read val
	fmt.Println(b)  //memory address
	fmt.Println(*b) //Value

	//change value with pointer
	*b = 10
	fmt.Println(a)

}
