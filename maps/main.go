package main

import "fmt"

func main() {
	//Define a map
	emails := make(map[string]string)
	names := map[string]string{"adam": "Adam Reuben"}

	fmt.Println(names)
	//Assign the map emails
	emails["bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Print(emails)
	fmt.Print(emails["bob"])
	fmt.Print(len(emails))

	//Delete map
	delete(emails, "bob") //delete --> comes with go

}
