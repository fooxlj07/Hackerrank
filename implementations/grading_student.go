package implementations

import (
	"fmt"
)

func GradingStudents(inputs []int) {
	for _, input := range inputs {
		fmt.Println(caculate(input))
	}
}

func caculate(grade int) int {
	if grade < 38 {
		return grade
	}
	multipleFive := grade / 5
	if (multipleFive+1)*5-grade >= 3 {
		grade = (multipleFive + 1) * 5
	}

	return grade
}
