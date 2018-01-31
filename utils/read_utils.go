package utils

//This file is utils which read input from console

import (
	"fmt"
)

//ReadIntsInput ...
//Read several lines of int,each line contain just one int
//first line idicate how many records it will receive
func ReadIntsInput(num int) []int {
	inputs := make([]int, num)

	for i := range inputs { // break the loop if text == "q"
		_, _ = fmt.Scanf("%d", &inputs[i])
	}
	return inputs
}

//ReadStringsInput ...
//read string seperated by space
func ReadStringInput() string {
	var str string
	_, _ = fmt.Scanln(&str)
	return str
}
