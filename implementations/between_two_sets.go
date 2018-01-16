package implementations

import (
	"fmt"

	"github.com/hacker_rank/utils"
)

func BetweenTwoSets() {
	nums := utils.ReadIntsInput(2)
	arrayA := utils.ReadIntsInput(nums[0])
	arrayB := utils.ReadIntsInput(nums[1])

	fmt.Println(nums, arrayA, arrayB)
	start := utils.Max(arrayA)
	end := utils.Min(arrayB)
	num := 0
	for i := start; i <= end; i++ {
		if isDivide(i, arrayA) && isDivided(arrayB, i) {
			num++
		}
	}
	fmt.Println(num)
}

func isDivide(num int, dividors []int) bool {
	for _, d := range dividors {
		if num%d != 0 {
			return false
		}
	}
	return true
}

func isDivided(nums []int, dividor int) bool {
	for _, n := range nums {
		if n%dividor != 0 {
			return false
		}
	}
	return true
}
