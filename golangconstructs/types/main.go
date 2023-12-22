package main

import (
	"fmt"
)

type Celcius float64
type Fahrenheit float64

func main() {
	var c Celcius = 100
	f := CtoF(c)
	fmt.Printf("%v celcius is %v fahrenheit", c, f)
	// FtoC(c) // Would throw even though the underlying types are the same.
}

func CtoF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}
