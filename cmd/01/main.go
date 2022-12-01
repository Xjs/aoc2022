package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	solution, err := aoc1task1(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
}

// aoc1task reads the input, splits double newlines, and sums up the numbers in the individual lines in all chunks,
// and returns the largest of the sums.
func aoc1task1(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)

	var sums []int
	var largestIdx int

	for s.Scan() {
		if len(sums) == 0 {
			sums = append(sums, 0)
		}

		current := len(sums) - 1

		t := s.Text()
		if t == "" {
			sums = append(sums, 0)
			continue
		}
		i, err := strconv.Atoi(t)
		if err != nil {
			return sums[largestIdx], err
		}

		sums[current] += i
		if sums[current] > sums[largestIdx] {
			largestIdx = current
		}
	}

	return sums[largestIdx], s.Err()
}
