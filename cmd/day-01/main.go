package main

import (
	"aoc2021/internal/pkg/utils"
	"fmt"
	"strconv"
)

const (
	fileTest = "./input.txt"
	fileOne  = "./input-1.txt"
	fileTwo  = "./input-2.txt"
)

func countIncreasingSums(filename string, windowSize int) {
	done := make(chan struct{})
	defer close(done)

	window := make([]int, windowSize)

	sumWindow := func(window []int) (sum int) {
		for _, v := range window {
			sum += v
		}
		return sum
	}

	count := -windowSize
	prevSum := 0

	stringsStream, streamErr := utils.ReadStrings(done, filename)
	for s := range stringsStream {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		window = append(window[1:], num)
		currSum := sumWindow(window)

		if prevSum < currSum {
			count++
		}

		prevSum = currSum
	}

	if err := <-streamErr; err != nil {
		panic(err)
	}

	fmt.Printf("The answer: %d\n", count)
}

func main() {
	countIncreasingSums(fileTest, 3)

	countIncreasingSums(fileOne, 1)
	countIncreasingSums(fileTwo, 3)
}
