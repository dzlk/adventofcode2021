package main

import (
	"aoc2021/internal/pkg/utils"
	"errors"
	"fmt"
	"strconv"
)

func countBits(filename string) {
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

func separateNumbersByBit(strings []string, bitIndex int) ([]string, []string) {
	zero := make([]string, 0)
	one := make([]string, 0)

	for _, s := range strings {
		if s[bitIndex] == '0' {
			zero = append(zero, s)
		} else {
			one = append(one, s)
		}
	}

	return zero, one
}

func findRating(numbers []string, criteria func(lhs int, rhs int) bool) (string, error) {
	if len(numbers) < 1 {
		return "", errors.New("too few numbers")
	}

	countBits := len(numbers[0])
	for bitIndex := 0; bitIndex < countBits && len(numbers) > 1; bitIndex++ {
		zero, one := separateNumbersByBit(numbers, bitIndex)

		if criteria(len(zero), len(one)) {
			numbers = zero
		} else {
			numbers = one
		}
	}

	if len(numbers) != 1 {
		return "", errors.New(fmt.Sprintf("left %d numbers", len(numbers)))
	}

	return numbers[0], nil
}

func verifyLifeSupportRating(filename string) {
	done := make(chan struct{})
	defer close(done)

	numbers, err := utils.ReadAllStrings(filename)
	if err != nil {
		panic(err)
	}

	oxygenRating, err := findRating(numbers, func(lhs int, rhs int) bool { return lhs > rhs })
	if err != nil {
		panic(err)
	}

	co2Rating, err := findRating(numbers, func(lhs int, rhs int) bool { return lhs <= rhs })
	if err != nil {
		panic(err)
	}

	fmt.Println("oxygen rating =", oxygenRating)
	fmt.Println("co2 rating =", co2Rating)

	oxygen, _ := strconv.ParseInt(oxygenRating, 2, 64)
	co2, _ := strconv.ParseInt(co2Rating, 2, 64)

	fmt.Printf("Answer: %d\n", oxygen*co2)

}

func main() {
	countBits("./input.txt")
	countBits("./input-1.txt")

	verifyLifeSupportRating("./input.txt")
	verifyLifeSupportRating("./input-2.txt")
}
