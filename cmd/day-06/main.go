package main

import (
	"aoc2021/internal/pkg/utils"
	"fmt"
)

func countFish(fish []int, days int) {
	day := 0

	counters := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, v := range fish {
		counters[v]++
	}

	for day < days {
		day++

		saving := counters[0]
		for i := len(counters) - 1; -1 < i; i-- {
			curr := counters[i]
			counters[i] = saving

			saving = curr
		}

		counters[6] += counters[8]

		//fmt.Printf("Day %d: %v\n", day, fish)
	}

	sum := 0
	for _, f := range counters {
		sum += f
	}

	fmt.Printf("Answer (%d): %d \n", days, sum)
}

func main() {
	fish, err := utils.GetNumbersFromString("3,4,3,1,2", ",")
	if err != nil {
		panic(err)
	}

	countFish(fish, 80)
	countFish(fish, 256)

	// ##
	fish, err = utils.GetNumbersFromString(
		"4,3,3,5,4,1,2,1,3,1,1,1,1,1,2,4,1,3,3,1,1,1,1,2,3,1,1,1,4,1,1,2,1,2,2,1,1,1,1,1,5,1,1,2,1,1,1,1,1,1,1,"+
			"1,1,3,1,1,1,1,1,1,1,1,5,1,4,2,1,1,2,1,3,1,1,2,2,1,1,1,1,1,1,1,1,1,1,4,1,3,2,2,3,1,1,1,4,1,1,1,1,5,1,1,"+
			"1,5,1,1,3,1,1,2,4,1,1,3,2,4,1,1,1,1,1,5,5,1,1,1,1,1,1,4,1,1,1,3,2,1,1,5,1,1,1,1,1,1,1,5,4,1,5,1,3,4,1,"+
			"1,1,1,2,1,2,1,1,1,2,2,1,2,3,5,1,1,1,1,3,5,1,1,1,2,1,1,4,1,1,5,1,4,1,2,1,3,1,5,1,4,3,1,3,2,1,1,1,2,2,1,"+
			"1,1,1,4,5,1,1,1,1,1,3,1,3,4,1,1,4,1,1,3,1,3,1,1,4,5,4,3,2,5,1,1,1,1,1,1,2,1,5,2,5,3,1,1,1,1,1,3,1,1,1,"+
			"1,5,1,2,1,2,1,1,1,1,2,1,1,1,1,1,1,1,3,3,1,1,5,1,3,5,5,1,1,1,2,1,2,1,5,1,1,1,1,2,1,1,1,2,1", ",")
	if err != nil {
		panic(err)
	}

	countFish(fish, 80)
	countFish(fish, 256)

}
