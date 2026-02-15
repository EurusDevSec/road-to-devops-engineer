package main

import "fmt"

func main() {
	// declare variables
	var name string
	var is_muggle bool
	var yearOld int
	fmt.Print("Enter your name & are you a muggle: ")
	// input parameter
	fmt.Scanf("%s %t %d", &name, &is_muggle, &yearOld)
	fmt.Println(name, is_muggle, yearOld)

}
