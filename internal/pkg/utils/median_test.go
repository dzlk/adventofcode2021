package utils

import "testing"

func assertMedian(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("Output %d not equal to expected %d\n", actual, expected)
	}
}

func TestFindMedianInEmpty(t *testing.T) {
	assertMedian(t, FindMedian([]int{}), 0)
}

func TestFindMedianInSingle(t *testing.T) {
	assertMedian(t, FindMedian([]int{10}), 10)
}

func TestMedianInUniqEven(t *testing.T) {
	assertMedian(t, FindMedian([]int{3, 2, 4, 1}), 2)
}

func TestMedianInUniqOdd(t *testing.T) {
	assertMedian(t, FindMedian([]int{3, 2, 4, 1, 5}), 3)
}

func TestFindMedianEven(t *testing.T) {
	assertMedian(t, FindMedian([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}), 2)
}

func TestFindMedianOdd(t *testing.T) {
	assertMedian(t, FindMedian([]int{16, 0, 1, 2, 0, 4, 2, 7, 1, 2, 14}), 2)
}
