package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var usePart2 bool
	flag.BoolVar(&usePart2, "part2", usePart2, "Use part 2 implementation")
	flag.Parse()

	f := task
	result, err := f(os.Stdin, usePart2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

type assignment struct{ left, right int }

func contains(a1, a2 assignment) bool {
	return a1.left <= a2.left && a1.right >= a2.right
}

func task(r io.Reader, part2 bool) (int, error) {
	var sum int

	s := bufio.NewScanner(r)
	for s.Scan() {
		assignments := strings.Split(s.Text(), ",")
		if len(assignments) != 2 {
			return 0, errors.New("need exactly two assignments per line")
		}

		var a [2]assignment

		for i, assignment := range assignments {
			numbers := strings.Split(assignment, "-")
			if len(numbers) != 2 {
				return 0, errors.New("need exactly two numbers per assignment")
			}
			var err error
			a[i].left, err = strconv.Atoi(numbers[0])
			if err != nil {
				return 0, err
			}
			a[i].right, err = strconv.Atoi(numbers[1])
			if err != nil {
				return 0, err
			}

			if !part2 {
				if contains(a[0], a[1]) || contains(a[1], a[0]) {
					sum++
				}
			} else {
				if !(a[0].right < a[1].left || a[1].right < a[0].left) {
					sum++
				}
			}
		}
	}

	return sum, s.Err()
}
