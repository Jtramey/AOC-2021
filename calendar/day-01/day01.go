package main

import (
	"advent-of-go-2021/utils/conv"
	"advent-of-go-2021/utils/files"
	"fmt"
)

func main() {
	input := conv.ToIntSlice(files.ReadFile(1, "\n"))
	fmt.Printf("Part 1: %d\n", solveDay1(input))
	fmt.Printf("Part 2: %d\n", solveDay1p2(input))
}

func solveDay1(input []int) int {
	count := 0
	prev := input[0]
	for i := 1; i < len(input); i++ {
		current := input[i]
		if current > prev {
			count++
		}
		prev = current
	}

	return count
}

func solveDay1p2(input []int) int {
	windows := make([]int, 0, len(input)/3)
	for i := 0; i < len(input)-2; i++ {
		windows = append(windows, input[i]+input[i+1]+input[i+2])
	}

	prev := windows[0]
	count := 0
	for i := 1; i < len(windows); i++ {
		if windows[i] > prev {
			count++
		}
		prev = windows[i]
	}

	return count
}
