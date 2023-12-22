package main

import (
	"fmt"
	"math/big"
)

func main() {
	s := "FF00EE"
	var decoded big.Int
	decoded.SetString(s[0:2], 16)
	fmt.Println(decoded.Int64())
	println(s[0:2])
}

func getColorComponents(s string) (int16, int16, int16) {
	return parseToInt16(s[0:2]), parseToInt16(s[2:4]), parseToInt16(s[4:6])
}

func parseToInt16(s string) int16 {
	var decoded big.Int
	decoded.SetString(s, 16)
	return int16(decoded.Int64())
}
