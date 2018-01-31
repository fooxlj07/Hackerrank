package sorting

import (
	"github.com/hacker_rank/utils"
)

func BigSorting() {
	var result []string
	num := utils.ReadIntsInput(1)
	str := utils.ReadStringInput()
	result = append(result, str)
	for i := 1; i < num[0]; i++ {
		str := utils.ReadStringInput()
		if len(str) > len(result[i-1]) {
			result = append(result, str)
		}
	}
}
