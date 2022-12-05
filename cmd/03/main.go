package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var usePart2 bool
	flag.BoolVar(&usePart2, "part2", usePart2, "Use part 2 implementation")
	flag.Parse()

	f := task1
	if usePart2 {
		f = task2
	}
	result, err := f(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func set(c string) map[rune]int {
	result := make(map[rune]int)
	for _, r := range c {
		result[r] += 1
	}
	return result
}

func matching(sn ...map[rune]int) rune {
	if len(sn) == 0 {
		return 0
	}
outer:
	for r, count := range sn[0] {
		if count == 0 {
			continue
		}
		for _, s2 := range sn[1:] {
			if s2[r] == 0 {
				continue outer
			}
		}
		return r
	}
	// should not happen
	return '\x00'
}

func score(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}
	if r >= 'A' && r <= 'Z' {
		return int(r) - int('A') + 27
	}
	return -1
}

func task1(r io.Reader) (int, error) {
	var sum int

	s := bufio.NewScanner(r)
	for s.Scan() {
		contents := s.Text()
		r, err := findMatching(contents)
		if err != nil {
			return 0, err
		}
		sum += score(r)
	}

	return sum, s.Err()
}

func task2(r io.Reader) (int, error) {
	var sum int
	var group []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		contents := s.Text()
		group = append(group, contents)
		if len(group) == 3 {
			var sets []map[rune]int
			for _, elf := range group {
				sets = append(sets, set(elf))
			}
			r := matching(sets...)
			sum += score(r)
			group = nil
		}
	}

	return sum, s.Err()
}

func findMatching(contents string) (rune, error) {
	if len(contents)%2 != 0 {
		return 0, errors.New("length of line must be divisible by 2")
	}

	c1, c2 := contents[:len(contents)/2], contents[len(contents)/2:]
	s1, s2 := set(c1), set(c2)
	r := matching(s1, s2)
	return r, nil
}
