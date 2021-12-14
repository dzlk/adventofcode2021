package utils

import (
	"math/rand"
	"time"
)

func FindMedian(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	idx := 0
	if len(nums)%2 == 0 {
		idx = len(nums)/2 - 1
	} else {
		idx = len(nums) / 2
	}

	return randomizeSelect(nums, idx)
}

func randomizeSelect(nums []int, idx int) int {
	if len(nums) == 1 {
		return nums[idx]
	}

	pivot, left, right := randomizePartition(nums)

	if idx < len(left) {
		return randomizeSelect(left, idx)
	}

	upperBound := len(nums) - len(right)
	if idx >= upperBound {
		return randomizeSelect(right, idx-upperBound)
	}

	return pivot
}

func randomizePartition(nums []int) (int, []int, []int) {
	pivot := SelectRandom(nums)

	left := make([]int, 0)
	right := make([]int, 0)

	for _, n := range nums {
		if n < pivot {
			left = append(left, n)
		}

		if n > pivot {
			right = append(right, n)
		}
	}

	return pivot, left, right
}

func SelectRandom(nums []int) int {
	rand.Seed(time.Now().UnixNano())

	return nums[rand.Intn(len(nums))]
}
