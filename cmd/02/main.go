package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Shape int

const (
	Unknown  Shape = 0
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3
)

type Battle int

const (
	UnknownBattle Battle = -1
	Lose          Battle = 0
	Draw          Battle = 3
	Win           Battle = 6
)

func parseShape(s rune) Shape {
	switch s {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	default:
		return Unknown
	}
}

func parseBattleOutcome(s rune) Battle {
	switch s {
	case 'X':
		return Lose
	case 'Y':
		return Draw
	case 'Z':
		return Win
	default:
		return UnknownBattle
	}
}

// battle returns 0 if the first shape wins,
// 3 if it's a draw and 6 if the second shape wins
func battle(a, b Shape) Battle {
	if a == b {
		return Draw
	}
	if (a == Rock && b == Scissors) ||
		(a == Scissors && b == Paper) ||
		(a == Paper && b == Rock) {
		return Lose
	}
	return Win
}

func parseBattle(s string) (int, error) {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return 0, errors.New("need exactly 2 parts")
	}
	if len(parts[0]) != 1 || len(parts[1]) != 1 {
		return 0, errors.New("need exactly 2 parts of length 1")
	}
	shapeA := parseShape(rune(parts[0][0]))
	shapeB := parseShape(rune(parts[1][0]))

	if shapeA == Unknown || shapeB == Unknown {
		return 0, fmt.Errorf("unknown shape in %q", s)
	}

	return int(shapeB) + int(battle(shapeA, shapeB)), nil
}

func parseBattleOutcomeSuggestion(s string) (int, error) {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return 0, errors.New("need exactly 2 parts")
	}
	if len(parts[0]) != 1 || len(parts[1]) != 1 {
		return 0, errors.New("need exactly 2 parts of length 1")
	}
	shapeA := parseShape(rune(parts[0][0]))
	outcome := parseBattleOutcome(rune(parts[1][0]))

	if shapeA == Unknown || outcome == UnknownBattle {
		return 0, fmt.Errorf("unknown letter in %q", s)
	}

	shapeB := needToPlay[shapeA][outcome]

	return int(shapeB) + int(battle(shapeA, shapeB)), nil
}

var needToPlay = map[Shape]map[Battle]Shape{
	Rock: {
		Lose: Scissors,
		Draw: Rock,
		Win:  Paper,
	},
	Paper: {
		Win:  Scissors,
		Lose: Rock,
		Draw: Paper,
	},
	Scissors: {
		Draw: Scissors,
		Win:  Rock,
		Lose: Paper,
	},
}

func main() {
	var usePart2 bool
	flag.BoolVar(&usePart2, "part2", usePart2, "Use part 2 implementation")
	flag.Parse()

	var r io.Reader = os.Stdin
	if usePart2 {
		fmt.Println(part2(r))
		return
	}
	fmt.Println(part1(r))
}

func part1(r io.Reader) int {
	var total int
	s := bufio.NewScanner(r)
	for s.Scan() {
		score, err := parseBattle(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		total += score
	}
	return total
}

func part2(r io.Reader) int {
	var total int
	s := bufio.NewScanner(r)
	for s.Scan() {
		score, err := parseBattleOutcomeSuggestion(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		total += score
	}
	return total
}
