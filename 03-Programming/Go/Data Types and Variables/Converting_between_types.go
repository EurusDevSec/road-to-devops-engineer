package main

import (
	"fmt"
	"strconv"
)

func main() {

	// var f float64 = 33.53
	// var i int = int(f)
	// fmt.Printf("%v\n", i)

	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("%v", s)

}
