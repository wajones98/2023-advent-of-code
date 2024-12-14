package day13

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 13

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
