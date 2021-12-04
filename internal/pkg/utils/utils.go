package utils

import (
	"bufio"
	"errors"
	"log"
	"os"
)

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
