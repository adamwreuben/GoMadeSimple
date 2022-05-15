package main

import "fmt"

func main() {
	fmt.Print("Hello world!")
	ids := []int{33, 45, 67, 45, 43, 63, 345}

	//loop through ids its like forEach() in js
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index i
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("SUm", sum)

	//Range with map
	names := map[string]string{"adam": "Adam Reuben"}

	for k, v := range names {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range names {
		fmt.Println("Names: " + k)
	}
}
