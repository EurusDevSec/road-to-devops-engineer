package main

import (
	"fmt"
	"reflect"
)

func main() {
	var grades int = 42
	var messeage string = "helloo world"

	fmt.Printf("variable grades=%v is of type %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("variable grades=%v is of type %v \n", messeage, reflect.TypeOf(messeage))

}
