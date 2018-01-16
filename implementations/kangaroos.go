package implementations

import (
	"fmt"
	"math"

	"github.com/hacker_rank/utils"
)

func Kangaroos() {
	inputs := utils.ReadIntsInput(4)
	distance := math.MaxInt16
	newDistance := inputs[2] - inputs[0]
	move1 := inputs[0]
	move2 := inputs[2]
	for newDistance < distance && newDistance != 0 {
		distance = newDistance
		move1 = move1 + inputs[1]
		move2 = move2 + inputs[3]
		newDistance = utils.Abs(move2 - move1)
	}
	if newDistance == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
