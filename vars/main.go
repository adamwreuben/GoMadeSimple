package main

import "fmt"

func main() {
	//int
	//string
	//bool
	//int int8 int16 int32
	//uint uint8 uint16 uint32 uintptr
	//byte -- alias for uint8
	//rune --- alias for int32
	//float32 float64
	//complex64 complex128

	//using var
	var name string = "Adam"
	const age = 23

	//name,age := "Adam",23
	//names := "Adam"
	fmt.Printf("%T\n", name)
}
