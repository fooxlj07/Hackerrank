package utils

import (
	"math"
)

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func Min(array []int) int {
	min := math.MaxInt16
	for _, v := range array {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(array []int) int {
	max := 0
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	return max
}
