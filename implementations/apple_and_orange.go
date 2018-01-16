package implementations

import (
	"fmt"

	"github.com/hacker_rank/utils"
)

func AppleAndOrange() {

	rangeLimit := utils.ReadIntsInput(2)

	positions := utils.ReadIntsInput(2)
	line3 := utils.ReadIntsInput(2)
	appleNum := line3[0]
	orangeNum := line3[1]

	appleDistances := utils.ReadIntsInput(appleNum)
	orangeDistances := utils.ReadIntsInput(orangeNum)
	fmt.Println(caculateFruit(positions[0], rangeLimit, appleDistances))
	fmt.Println(caculateFruit(positions[1], rangeLimit, orangeDistances))
}

func caculateFruit(pos int, rangeLimit, allDistaces []int) int {
	sum := 0
	for _, d := range allDistaces {
		currentPos := pos + d
		if currentPos >= rangeLimit[0] && currentPos <= rangeLimit[1] {
			sum = sum + 1
		}
	}
	return sum
}
