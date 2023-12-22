package main

import (
	"math"
	"slices"
)

func Min(values []float64) float64 {
	res, _ := Minmax(values)
	return res
}
func Max(values []float64) float64 {
	_, res := Minmax(values)
	return res
}

func Minmax(values []float64) (float64, float64) {
	values = removeNan(values)
	slices.Sort(values)
	return values[0], values[len(values)-1]
}

func removeNan(values []float64) []float64 {
	result := *new([]float64)
	for _, x := range values {
		if !math.IsNaN(x) {
			result = append(result, x)
		}
	}
	return result
}

func Avg(values []float64) float64 {
	return sum(values) / float64(len(values))
}

func sum(values []float64) float64 {
	var result float64
	for _, i := range values {
		result += i
	}
	return result
}
