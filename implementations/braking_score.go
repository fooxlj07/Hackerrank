package implementations

import (
	"fmt"

	"github.com/hacker_rank/utils"
)

func BreakingScore() {
	seasonNum := utils.ReadIntsInput(1)
	seasonScores := utils.ReadIntsInput(seasonNum[0])
	breackingScore := []int{0, 0}
	min := seasonScores[0]
	max := seasonScores[0]

	for i := 1; i <= seasonNum[0]; i++ {
		if seasonScores[i] < min {
			breackingScore[1] = breackingScore[1] + 1
			min = seasonScores[i]
		}

		if seasonScores[i] > max {
			breackingScore[0] = breackingScore[0] + 1
			max = seasonScores[i]
		}
	}
	fmt.Println(breackingScore)
}
