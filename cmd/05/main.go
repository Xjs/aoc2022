package main

import (
	"bufio"
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

	crates, err := task(os.Stdin, usePart2)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range crates {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

func task(r io.Reader, part2 bool) ([]crate, error) {
	s := bufio.NewScanner(r)
	stacks, err := parseStacks(s)
	if err != nil {
		return nil, err
	}
	instructions, err := parseInstructions(s)
	if err != nil {
		return nil, err
	}
	apply(stacks, instructions, part2)

	var crates []crate
	for _, stack := range stacks {
		c, _ := stack.pop()
		crates = append(crates, c)
	}
	return crates, nil
}

type crate rune

type stack []crate

func (s stack) push(c ...crate) stack {
	return append(s, c...)
}

func (s stack) pop() (crate, stack) {
	c, s := s.popN(1)
	return c[0], s
}

func (s stack) popN(n int) ([]crate, stack) {
	return s[len(s)-n:], s[:len(s)-n]
}

type instruction struct {
	quantity int
	from     int
	to       int
}

func parseInstruction(s string) (instruction, error) {
	words := strings.Fields(s)
	if len(words) != 6 {
		return instruction{}, fmt.Errorf("expected instruction like 'move <n> from <a> to <b>, got %q'", s)
	}
	if words[0] != "move" {
		return instruction{}, fmt.Errorf("unknown verb %q", words[0])
	}
	if words[2] != "from" || words[4] != "to" {
		return instruction{}, fmt.Errorf("unknown directionality in %q", s)
	}

	quantity, err := strconv.Atoi(words[1])
	if err != nil {
		return instruction{}, fmt.Errorf("error parsing quantity %q: %w", words[1], err)
	}

	from, err := strconv.Atoi(words[3])
	if err != nil {
		return instruction{}, fmt.Errorf("error parsing from-index %q: %w", words[3], err)
	}

	to, err := strconv.Atoi(words[5])
	if err != nil {
		return instruction{}, fmt.Errorf("error parsing to-index %q: %w", words[5], err)
	}

	// correct indexes here (elves start counting at 1)
	return instruction{quantity: quantity, from: from - 1, to: to - 1}, nil
}

func parseInstructions(s *bufio.Scanner) ([]instruction, error) {
	var instructions []instruction
	for s.Scan() {
		i, err := parseInstruction(s.Text())
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, i)
	}
	return instructions, s.Err()
}

// parseStacks parses only the stacks from the given scanner.
// The scanner is expected to be reused for the instructions.
func parseStacks(s *bufio.Scanner) ([]stack, error) {
	var lines [][]crate
	var n int

	const width = 4

scan:
	for s.Scan() {
		l := s.Text()

		if l == "" {
			break
		}

		if n == 0 {
			// hardcoded width of columns
			n = (len(s.Text()) + 1) / width
		} else if len(l) != n*width-1 {
			return nil, fmt.Errorf("invalid line width %d, want %d", len(l), n*width)
		}

		line := make([]crate, n)

		for i := 0; i < n; i++ {
			cRaw := l[i*4 : ((i+1)*4)-1]
			c := strings.Trim(cRaw, "[] ")
			if len(c) > 1 {
				return nil, fmt.Errorf("invalid crate %q", c)
			}

			if len(c) == 0 {
				line[i] = 0
				continue
			}

			if !('A' <= c[0] && c[0] <= 'Z') {
				// we have reached the numbers row
				continue scan
			}

			line[i] = crate(c[0])
		}

		lines = append(lines, line)
	}

	stacks := make([]stack, n)
	for row := len(lines) - 1; row >= 0; row-- {
		for col := 0; col < n; col++ {
			c := lines[row][col]
			if c == 0 {
				continue
			}
			stacks[col] = stacks[col].push(c)
		}
	}

	return stacks, nil
}

func applySingle(stacks []stack, inst instruction, part2 bool) {
	if part2 {
		var c []crate
		c, stacks[inst.from] = stacks[inst.from].popN(inst.quantity)
		stacks[inst.to] = stacks[inst.to].push(c...)
		return
	}

	for ; inst.quantity > 0; inst.quantity-- {
		var c crate
		c, stacks[inst.from] = stacks[inst.from].pop()
		stacks[inst.to] = stacks[inst.to].push(c)
	}
}

func apply(stacks []stack, instructions []instruction, part2 bool) {
	for _, instruction := range instructions {
		applySingle(stacks, instruction, part2)
	}
}
