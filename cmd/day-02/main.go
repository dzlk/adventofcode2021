package main

import (
	"aoc2021/internal/pkg/utils"
	"fmt"
	"strconv"
	"strings"
)

func detectPosition(filename string) {
	done := make(chan struct{})
	defer close(done)

	withoutAim := struct {
		position int
		depth    int
	}{0, 0}
	withAim := struct {
		aim      int
		position int
		depth    int
	}{0, 0, 0}

	stringsStream, streamErr := utils.ReadStrings(done, filename)
	for s := range stringsStream {
		parts := strings.Fields(s)

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		if parts[0] == "forward" {
			withoutAim.position += value

			withAim.position += value
			withAim.depth += withAim.aim * value
		}

		if parts[0] == "down" {
			withoutAim.depth += value

			withAim.aim += value
		}

		if parts[0] == "up" {
			withoutAim.depth -= value

			withAim.aim -= value
		}
	}

	if err := <-streamErr; err != nil {
		panic(err)
	}

	fmt.Printf("WithoutAim. Horizontal position: %d, depth: %d. The answer: %d\n",
		withoutAim.position, withoutAim.depth, withoutAim.position*withoutAim.depth)

	fmt.Printf("WithAim. Horizontal position: %d, depth: %d. The answer: %d\n\n",
		withAim.position, withAim.depth, withAim.position*withAim.depth)
}

func main() {
	detectPosition("./input.txt")
	detectPosition("./input-1.txt")
	detectPosition("./input-2.txt")
}
