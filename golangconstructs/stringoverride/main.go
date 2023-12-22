package main

import (
	"fmt"
)

type Widget float32

func (w Widget) String() string { return fmt.Sprintf("%v widgets", float32(w)) }
func main() {
	var x Widget = 12
	fmt.Println(x)
}
