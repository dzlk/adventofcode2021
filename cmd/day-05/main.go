package main

import (
	"aoc2021/internal/pkg/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func getCoordinates(str string) (int, int, error) {
	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("wrong format")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

type Report struct {
	x1, y1, x2, y2 int
}

func parseReport(line string) (*Report, error) {
	coords := strings.Split(line, " -> ")
	if len(coords) != 2 {
		return nil, errors.New("wrong format")
	}

	x1, y1, err := getCoordinates(coords[0])
	if err != nil {
		return nil, err
	}

	x2, y2, err := getCoordinates(coords[1])
	if err != nil {
		return nil, err
	}

	return &Report{x1, y1, x2, y2}, nil
}

func generateDangerZones(report *Report, withoutDiagonals bool) <-chan string {
	zones := make(chan string)

	go func() {
		defer close(zones)

		startX := report.x1
		endX := report.x2
		changeX := 1
		if startX > endX {
			changeX = -1
		}

		startY := report.y1
		endY := report.y2
		changeY := 1
		if startY > endY {
			changeY = -1
		}

		if withoutDiagonals && startX != endX && startY != endY {
			return
		}

		x := startX
		y := startY
		for x != endX || y != endY {
			//fmt.Printf("(x, y) = (%d, %d) => %d\n", x, y, x*K+y)
			zones <- serializeXY(x, y)

			if x != endX {
				x += changeX
			}

			if y != endY {
				y += changeY
			}
		}
		//fmt.Printf("(x, y) = (%d, %d) => %d\n", x, y, x*K+y)
		zones <- serializeXY(x, y)
	}()

	return zones
}

func serializeXY(x int, y int) string {
	return fmt.Sprintf("(%d,%d)", x, y)
}

func findDangersZone(filename string, withoutDiagonals bool) {
	done := make(chan struct{})
	defer close(done)

	marks := make(map[string]int)
	dangersZones := 0

	stringsStream, streamErr := utils.ReadStrings(done, filename)
	for str := range stringsStream {
		report, err := parseReport(str)
		if err != nil {
			panic(err)
		}

		//fmt.Println(report)
		for key := range generateDangerZones(report, withoutDiagonals) {
			marks[key]++
			if marks[key] == 2 {
				dangersZones++
			}
		}
	}

	if err := <-streamErr; err != nil {
		panic(err)
	}

	//fmt.Println(marks)
	fmt.Println("Answer", dangersZones)

}

func main() {
	findDangersZone("./input.txt", true)
	findDangersZone("./input-1.txt", true)

	findDangersZone("./input.txt", false)
	findDangersZone("./input-1.txt", false)
}
