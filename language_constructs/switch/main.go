package main

import (
	"fmt"
)

func main() {
	items := []int{-12, -1, 0, 1, 2, 3, 4}

	for _, i := range items {
		result := get_switch_value_correct(i)
		fmt.Printf("The result of passing %d to get_switch_value is %q\n", i, result)
	}

}

// This will work, but the line with the default statement will complain that
// it should be first or last.
// func get_switch_value(x int) string {
// 	switch {
// 	case x < 0:
// 		return "Less than zero"
// 	default: // note the value is out of order
// 		return "Equal to zero"
// 	case x > 0:
// 		return "Greater than zero"
// 	}
// }

func get_switch_value_correct(x int) string {
	switch {
	case x < 0:
		return "Less than zero"
	case x > 0:
		return "Greater than zero"
	default: // note the value is out of order
		return "Equal to zero"
	}

}
