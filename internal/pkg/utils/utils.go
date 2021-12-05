package utils

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetNumbersFromString(str string, sep string) ([]int, error) {
	nums := make([]int, 0)

	for _, p := range strings.Split(str, sep) {
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return nums, err
		}

		nums = append(nums, n)
	}

	return nums, nil
}

func ReadAllStrings(filename string) ([]string, error) {
	done := make(chan struct{})
	defer close(done)

	strs := make([]string, 0)

	stringsStream, streamErr := ReadStrings(done, filename)
	for s := range stringsStream {
		strs = append(strs, s)
	}

	if err := <-streamErr; err != nil {
		return nil, err
	}

	return strs, nil
}

func ReadStrings(done <-chan struct{}, filename string) (<-chan string, <-chan error) {
	out := make(chan string)
	outErr := make(chan error, 1)

	go func() {
		defer close(out)

		outErr <- func() error {
			file, err := os.Open(filename)
			if err != nil {
				return err
			}

			defer func() {
				if err = file.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() { // internally, it advances token based on sperator
				select {
				case out <- scanner.Text():
				case <-done:
					return errors.New("walk cancelled")
				}
			}

			return nil
		}()
	}()

	return out, outErr
}
