package main

import (
	"aoc2021/internal/pkg/utils"
	"fmt"
	"strconv"
)

func countBytes(filename string) {
	done := make(chan struct{})
	defer close(done)

	countBits := make([]int, 0)
	count := 0

	stringsStream, streamErr := utils.ReadStrings(done, filename)
	for s := range stringsStream {
		if len(countBits) == 0 {
			countBits = make([]int, len(s))
		}

		for i, r := range s {
			if r == '1' {
				countBits[i]++
			}
		}

		count++
	}

	if err := <-streamErr; err != nil {
		panic(err)
	}

	count /= 2

	min := ""
	max := ""

	for _, x := range countBits {
		if x > count {
			max += "1"
			min += "0"
		} else {
			max += "0"
			min += "1"
		}
	}

	minDec, _ := strconv.ParseInt(min, 2, 64)
	maxDec, _ := strconv.ParseInt(max, 2, 64)

	fmt.Printf("Min = %s (%d); Max = %s (%d). Answer = %d * %d = %d\n",
		min, minDec, max, maxDec,
		minDec, maxDec, minDec*maxDec)
}

func main() {
	countBytes("./input.txt")
	countBytes("./input-1.txt")
}
