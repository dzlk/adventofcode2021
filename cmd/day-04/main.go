package main

import (
	"aoc2021/internal/pkg/utils"
	"errors"
	"fmt"
)

type Game struct {
	boards  []*Board
	numbers []int
}

func NewGame(numbers []int) *Game {
	return &Game{
		make([]*Board, 0),
		numbers,
	}
}

func (g *Game) AddBoard(b *Board) error {
	if len(g.boards) > 0 && g.boards[0].Capacity() != b.Capacity() {
		return errors.New("board has different number count")
	}

	g.boards = append(g.boards, b)
	return nil
}

func (g *Game) Play() (*Board, int) {
	for _, num := range g.numbers {
		for _, b := range g.boards {
			if b.CheckNumAndCheckWin(num) {
				return b, num
			}
		}
	}

	return nil, 0
}

const (
	coefCol   = 100
	coefIndex = 100000
)

type Board struct {
	/*
	  2 2

	  0 1 2 3 4
	  5 6 7 8 9
	*/
	mappingNums map[int]int
	nums        []int
	checks      map[int]int
	cols        int
	winCol      int
	winRow      int
}

func NewBoard() *Board {
	return &Board{
		make(map[int]int),
		make([]int, 0),
		make(map[int]int),
		0,
		-1,
		-1,
	}
}

func (b *Board) Capacity() int {
	return len(b.nums)
}

func (b *Board) AddLine(nums []int) error {
	if b.cols > 0 && b.cols != len(nums) {
		return errors.New("rows has different length")
	}

	b.cols = len(nums)

	for _, n := range nums {
		b.mappingNums[n] = len(b.nums)
		b.nums = append(b.nums, n)
	}

	return nil
}

func (b *Board) CheckNumAndCheckWin(num int) bool {
	index, found := b.mappingNums[num]

	if found {
		r := index / b.cols
		c := index % b.cols

		ri := r
		ci := c + coefCol

		b.checks[ri]++
		b.checks[ci]++
		b.checks[index+coefIndex] = 1

		if b.checks[ri] == b.cols {
			b.winRow = r
			return true
		}

		if b.checks[ci] == b.cols {
			b.winCol = c
			return true
		}
	}

	return false
}

func (b *Board) getWinnerIndexes() []int {
	indexes := make([]int, 0)

	if b.winRow > -1 {
		for i := 0; i < b.cols; i++ {
			indexes = append(indexes, b.winRow+i)
		}
	} else if b.winCol > -1 {
		for i := 0; i < b.cols; i++ {
			indexes = append(indexes, b.winCol+i*5)
		}
	}

	return indexes
}

func (b *Board) getWinnerNumbers() []int {
	nums := make([]int, 0)

	for _, i := range b.getWinnerIndexes() {
		nums = append(nums, b.nums[i])
	}

	return nums
}

func (b *Board) getSumUncheckedNumbers() int {
	sum := 0

	for i, n := range b.nums {
		if _, ok := b.checks[i+coefIndex]; !ok {
			fmt.Print(n, " ")
			sum += n
		}
	}
	fmt.Println()

	return sum
}

func readGame(filename string) (*Game, error) {
	done := make(chan struct{})
	defer close(done)

	var game *Game
	var board *Board

	stringsStream, streamErr := utils.ReadStrings(done, filename)
	for str := range stringsStream {
		if game == nil {
			nums, err := utils.GetNumbersFromString(str, ",")
			if err != nil {
				return game, err
			}

			game = NewGame(nums)
			continue
		}

		if str == "" && board == nil {
			continue
		}

		if str == "" {
			err := game.AddBoard(board)
			if err != nil {
				return game, nil
			}

			board = nil
			continue
		}

		if board == nil {
			board = NewBoard()
		}

		nums, err := utils.GetNumbersFromString(str, " ")
		if err != nil {
			return game, err
		}
		err = board.AddLine(nums)
		if err != nil {
			return game, err
		}
	}

	if board != nil {
		err := game.AddBoard(board)
		if err != nil {
			return game, nil
		}
	}

	if err := <-streamErr; err != nil {
		return game, err
	}

	return game, nil
}

func playGame(filename string) {
	game, err := readGame(filename)
	if err != nil {
		panic(err)
	}

	board, num := game.Play()
	fmt.Printf("%d\n%v\n", num, board)
	winnersNums := board.getWinnerNumbers()
	fmt.Println(winnersNums)

	sum := board.getSumUncheckedNumbers()
	fmt.Printf("Sum: %d, Answer: %d", sum, sum*num)
}

func main() {
	playGame("./input.txt")
	fmt.Println("------")
	playGame("./input-1.txt")
}
