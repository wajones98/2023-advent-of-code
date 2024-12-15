package day13

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const (
	Day        int = 13
	ATokenCost     = 3
	BTokenCost     = 1
)

func Run() (*days.Result[int, int], error) {
	pOne, err := Part1()
	if err != nil {
		return nil, err
	}

	pTwo, err := Part2()
	if err != nil {
		return nil, err
	}

	return &days.Result[int, int]{
		Part1: pOne,
		Part2: pTwo,
	}, nil
}

func Part1() (int, error) {
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	prizes, err := LoadInput(s)
	if err != nil {
		return 0, err
	}

	return TotalTokens(prizes, 0), nil
}

func Part2() (int, error) {
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	return 0, nil
}

type Coords struct {
	X, Y int
}

type Prize struct {
	Location,
	ButtonA,
	ButtonB Coords
}

func LoadInput(s *bufio.Scanner) ([]Prize, error) {
	prizes := []Prize{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}

		buttonA := ParseLine(line, "+")
		s.Scan()
		buttonB := ParseLine(s.Text(), "+")
		s.Scan()
		prizeLocation := ParseLine(s.Text(), "=")

		prizes = append(prizes, Prize{
			Location: prizeLocation,
			ButtonA:  buttonA,
			ButtonB:  buttonB,
		})
	}
	return prizes, nil
}

func ParseLine(line, delimeter string) Coords {
	parts := strings.Split(line, ":")
	parts = strings.Split(parts[1], ",")

	x, _ := strconv.Atoi(strings.Split(parts[0], delimeter)[1])
	y, _ := strconv.Atoi(strings.Split(parts[1], delimeter)[1])

	return Coords{
		X: x,
		Y: y,
	}
}

func PossibleCombinations(location, x, y int) (map[int]int, bool) {
	combinations := map[int]int{}
	maxIterations := location / gcd(x, y) * 2
	for a := range maxIterations {
		b := (location - x*a) / y
		if b < 0 || a*x+b*y != location {
			continue
		}
		combinations[a] = int(b)
	}
	return combinations, len(combinations) > 0
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func FindCheapestCombination(xCombinations, yCombinations map[int]int) int {
	cheapest := math.MaxInt

	for a, b := range xCombinations {
		if yb, ok := yCombinations[a]; ok && yb == b {
			cost := (a * ATokenCost) + (b * BTokenCost)
			if cost < cheapest {
				cheapest = cost
			}
		}
	}
	if cheapest == math.MaxInt {
		return 0
	}
	return cheapest
}

func FindTokenCost(prize Prize, offset int) int {

	return 0
}

func TotalTokens(prizes []Prize, offset int) int {
	total := 0
	for _, prize := range prizes {
		total += FindTokenCost(prize, offset)
	}

	return total
}
