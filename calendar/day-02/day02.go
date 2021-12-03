package main

import (
	"advent-of-go-2021/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(2, "\n")
	fmt.Printf("Part 1: %d\n", solveDay2(input))
}

func solveDay2(input []string) int64 {
	x := int64(0)
	y := int64(0)
	aim := int64(0)
	for i := range input {
		split := strings.SplitAfter(input[i], " ")
		dist, _ := strconv.ParseInt(split[1], 10, 0)
		dir := strings.TrimSpace(split[0])
		switch dir {
		case "forward":
			x += dist
			y += aim * dist
		case "up":
			aim -= dist
		case "down":
			aim += dist
		default:
			continue
		}
	}
	return x * y
}
