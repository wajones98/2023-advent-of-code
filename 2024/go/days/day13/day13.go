package day13

import (
	"bufio"
	"fmt"
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
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	return 0, nil
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
	for a := range 200 {
		if math.Mod(float64(location)-(float64(x)*float64(a)), float64(y)) != 0 {
			continue
		}

		b := (float64(location) - (float64(x) * float64(a))) / float64(y)
		if b <= 0 {
			continue
		}
		combinations[a] = int(b)
	}
	return combinations, len(combinations) > 0
}

func FindCheapestCombination(xCombinations, yCombinations map[int]int) int {
	cheapest := math.MaxInt

	for a, b := range xCombinations {
		if _, ok := yCombinations[a]; ok {
			cost := (a * ATokenCost) + (b * BTokenCost)
			fmt.Printf("A: %d, B: %d = %d\n", a, b, cost)
			if cost < cheapest {
				cheapest = cost
			}
		}
	}

	return cheapest
}

func FindTokenCost(prize Prize) int {
	xCombinations, ok := PossibleCombinations(prize.Location.X, prize.ButtonA.X, prize.ButtonB.X)
	if !ok {
		return 0
	}

	yCombinations, ok := PossibleCombinations(prize.Location.Y, prize.ButtonA.Y, prize.ButtonB.Y)
	if !ok {
		return 0
	}

	temp := FindCheapestCombination(xCombinations, yCombinations)
	fmt.Printf("%d\n", temp)
	return temp
}
