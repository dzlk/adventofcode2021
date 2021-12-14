package main

import (
	"aoc2021/internal/pkg/utils"
	"fmt"
)

func main() {
	calcFuel([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 0)
	fmt.Println("------")

	calcFuel([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 1)
	fmt.Println("------")

	nums, err := utils.ReadNumbers("./input-1.txt", ",")
	if err != nil {
		panic(err)
	}

	calcFuel(nums, 0)
	fmt.Println("------")
	calcFuel(nums, 1)
	fmt.Println("------")
}

func calcFuel(positions []int, cost int) {
	median := utils.FindMedian(positions)
	fmt.Println("Sync position: ", median)

	fuel := 0
	for _, pos := range positions {
		distance := pos - median
		if distance < 0 {
			distance *= -1
		}

		firstStepCost := 1
		lastStepCost := firstStepCost + (distance-1)*cost

		fuel += (firstStepCost + lastStepCost) * distance / 2
	}

	fmt.Println("Total spent fuel: ", fuel)
}
