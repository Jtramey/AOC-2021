package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// got tired of the off by one errors.
//go:embed puzzle-input.in
var input string

func main() {
	solveDay3()
}

func solveDay3() {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	var g, e string
	for i := 0; i < len(lines[0]); i++ {
		ones, zeros := count1and0(i, lines)
		if ones > zeros {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}
	nG, _ := strconv.ParseInt(g, 2, 64)
	nE, _ := strconv.ParseInt(e, 2, 64)
	fmt.Println(nG * nE)

	// Part2
	var ox string
	f := lines
	for i := 0; len(f) > 1; i++ {
		mp := mostPop(count1and0(i, f))
		var tmp []string
		for _, line := range f {
			if line[i] == mp {
				tmp = append(tmp, line)
			}
		}
		f = tmp
	}
	if len(f) > 0 {
		ox = f[0]
	}

	var co2 string
	f = lines
	for pos := 0; len(f) > 1; pos++ {
		lp := leastPop(count1and0(pos, f))
		var tmp []string
		for _, line := range f {
			if line[pos] == lp {
				tmp = append(tmp, line)
			}
		}
		f = tmp
	}
	if len(f) > 0 {
		co2 = f[0]
	}
	nox, _ := strconv.ParseInt(ox, 2, 64)
	nco2, _ := strconv.ParseInt(co2, 2, 64)
	fmt.Println(nox * nco2)
}

func count1and0(i int, input []string) (ones int, zeros int) {
	for _, line := range input {
		if line[i] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return
}

func leastPop(ones int, zeros int) byte {
	if ones < zeros {
		return '1'
	}
	return '0'
}

func mostPop(ones int, zeros int) byte {
	if ones >= zeros {
		return '1'
	}
	return '0'
}
