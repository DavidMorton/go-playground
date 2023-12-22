package main

import (
	"fmt"
	"reflect"
)

func main() {
	demo_new()
}

func demo_new() {
	p := new(int)
	t := reflect.TypeOf(p)
	fmt.Printf("The type of p is %v\n", t)
}
