package main

import (
	"fmt"
	"strconv"
)

func main() {

	// var f float64 = 33.53
	// var i int = int(f)
	// fmt.Printf("%v\n", i)

	//--------CONVERT INT -> STR----------
	// var i int = 42
	// var s string = strconv.Itoa(i)
	// fmt.Printf("%q", s)

	// -----CONVERT STR -> INT ---------
	var s string = "200"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T", err, err)

}
