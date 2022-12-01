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
	fn := aoc1task1
	if len(os.Args) > 1 && os.Args[1] == "part2" {
		fn = aoc1task2
	}

	solution, err := fn(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
}

func aoc1task1(r io.Reader) (int, error) {
	return aoc1(r, 1)
}

func aoc1task2(r io.Reader) (int, error) {
	return aoc1(r, 3)
}

// aoc1 reads the input, splits double newlines, and sums up the numbers in the individual lines in all chunks,
// and returns the sum of the count largest of the sums.
func aoc1(r io.Reader, count int) (int, error) {
	s := bufio.NewScanner(r)

	var sums []int
	largestIdxs := make([]int, count)
	for c := 0; c < count; c++ {
		largestIdxs[c] = -1
	}

	var inserted bool
	insert := func(current int) {
		if inserted {
			return
		}
		inserted = true
		for c := 0; c < count; c++ {
			idx := largestIdxs[c]
			if idx == -1 {
				largestIdxs[c] = current
				return
			}
			if sums[current] > sums[idx] {
				for cc := count - 1; cc > c; cc-- {
					largestIdxs[cc] = largestIdxs[cc-1]
				}
				largestIdxs[c] = current
				return
			}
		}
	}

	current := -1

	for s.Scan() {
		t := s.Text()
		if t == "" {
			if current == -1 {
				continue
			}

			sums = append(sums, 0)
			insert(current)
			continue
		}

		i, err := strconv.Atoi(t)
		if err != nil {
			return sums[largestIdxs[0]], err
		}

		inserted = false
		if len(sums) == 0 {
			sums = append(sums, 0)
		}
		current = len(sums) - 1

		sums[current] += i
	}

	if current != -1 {
		insert(current)
	}

	sum := 0
	for _, li := range largestIdxs {
		if li == -1 {
			continue
		}
		sum += sums[li]
	}

	return sum, s.Err()
}
