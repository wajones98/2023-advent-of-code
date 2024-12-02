package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/input"

	"github.com/wajones98/advent-of-code/days"
)

type Direction = bool

type DetermineDirectionError struct{}

func (e DetermineDirectionError) Error() string {
	return "Values are the same"
}

const (
	Day        int       = 2
	Increasing Direction = true
	Decreasing Direction = false
)

var DirectionError DetermineDirectionError = DetermineDirectionError{}

func Run() (*days.Result[int, int], error) {

	reports, err := loadReports()
	if err != nil {
		return nil, err
	}

	log.Printf("%v\n", reports)

	return &days.Result[int, int]{
		Part1: Part1(),
		Part2: Part2(),
	}, nil
}

func loadReports() ([][]uint64, error) {
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return nil, err
	}
	defer closeFile()

	reports := [][]uint64{}

	for s.Scan() {
		line := s.Text()
		report := []uint64{}
		for _, val := range strings.Split(line, " ") {
			num, err := strconv.ParseUint(val, 10, 0)
			if err != nil {
				return nil, err
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func Part1() int {
	return 0
}

func determineDirection(left, right uint64) (Direction, error) {
	if left == right {
		return false, DirectionError
	}
	return Direction(left < right), nil
}

func isSafe(left, right uint64, direction Direction) bool {
	if left == right {
		return false
	}

	var diff uint64 = 0
	if direction {
		if left > right {
			return false
		}
		diff = right - left
	} else {
		if left < right {
			return false
		}
		diff = left - right
	}

	if diff > 3 {
		return false
	}

	return true
}

func Part2() int {
	return 0
}
