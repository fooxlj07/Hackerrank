package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/hacker_rank/f"
)

func main() {
	var inputs []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numStr := scanner.Text()
	num, _ := strconv.Atoi(numStr)
	for i := 0; i < num; i++ { // break the loop if text == "q"
		scanner.Scan()
		text := scanner.Text()
		n, _ := strconv.Atoi(text)
		inputs = append(inputs, n)
	}
	f.GradingStudents(inputs)
}
