package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const Head = 1
const Tail = 2

type Motion struct {
	X, Y int
}

// A Point represents a point on a
type Point struct {
	X, Y int
}

// P is a convenience constructor for Point.
func P(x, y int) Point {
	return Point{X: x, Y: y}
}

var (
	Down  = Motion{0, -1}
	Up    = Motion{0, 1}
	Left  = Motion{-1, 0}
	Right = Motion{1, 0}
)

func parseMotion(s string) (Motion, error) {
	switch s {
	case "R":
		return Right, nil
	case "U":
		return Up, nil
	case "L":
		return Left, nil
	case "D":
		return Down, nil
	default:
		return Motion{}, fmt.Errorf("unknown motion %q", s)
	}
}

func (m Motion) Apply(p Point) Point {
	x := p.X + m.X
	y := p.Y + m.Y

	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

	p.X = x
	p.Y = y

	return p
}

func main() {
	head := P(0, 0)
	tail := P(0, 0)

	visited := make(map[Point]struct{})
	visited[tail] = struct{}{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		words := strings.Fields(s.Text())
		if len(words) != 2 {
			log.Fatalf("need 2 words per line, got %q", line)
		}
		m, err := parseMotion(words[0])
		if err != nil {
			log.Fatalf("error parsing motion: %v", err)
		}
		steps, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatalf("error parsing steps: %v", err)
		}

		for ; steps > 0; steps-- {
			head = m.Apply(head)
			tm := follow(head, tail)
			tail = tm.Apply(tail)
			visited[tail] = struct{}{}
		}
	}

	fmt.Println(len(visited))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func norm(x int) int {
	if x == 0 {
		return x
	}
	return x / abs(x)
}

func follow(head, tail Point) Motion {
	var result Motion

	dx, dy := head.X-tail.X, head.Y-tail.Y
	adx := abs(dx)
	ady := abs(dy)

	if adx > 1 || ady > 1 {
		result.X = norm(dx)
		result.Y = norm(dy)
	}

	return result
}
