package main

import "fmt"

func main() {
	x, y := 5, 10
	if x < y {

		fmt.Printf("%d is less than %d\n", x, y)

	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//switch
	color := "red"
	switch color {
	case "red":
		fmt.Print("Color is red")
	case "blue":
		fmt.Print("Color is blue")
	default:
		fmt.Print("Color is NOT blue or Red")
	}
}
