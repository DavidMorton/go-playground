package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	var x uint64 = 1088444
	var result int

	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	fmt.Println(pc)
	fmt.Println(result)
}
