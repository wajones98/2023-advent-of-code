package day2

import (
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/input"

	"github.com/wajones98/advent-of-code/days"
)

func Run(day int) (*days.Result[int, int], error) {
	s, closeFile, err := input.GetInput(day)
	if err != nil {
		return nil, err
	}
	defer closeFile()

	reports := []uint64{}

	for s.Scan() {
		line := s.Text()
		for _, val := range strings.Split(line, " ") {
			num, err := strconv.ParseUint(val, 10, 0)
			if err != nil {
				return nil, err
			}
			reports = append(reports, num)
		}
	}

	return &days.Result[int, int]{
		Part1: Part1(),
		Part2: Part2(),
	}, nil
}

func Part1() int {
	return 0
}

func Part2() int {
	return 0
}
