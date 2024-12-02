package day2

import (
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

	return &days.Result[int, int]{
		Part1: Part1(reports),
		Part2: Part2(reports),
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

func Part1(reports [][]uint64) int {
	total := 0
	for _, report := range reports {
		if reportIsSafe(report) {
			total += 1
		}
	}
	return total
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

func reportIsSafe(report []uint64) bool {
	var direction Direction
	for i := 1; i < len(report); i++ {
		left := report[i-1]
		right := report[i]
		if i == 1 {
			d, err := determineDirection(left, right)
			direction = d
			if err != nil {
				return false
			}
		}

		ok := isSafe(left, right, direction)
		if !ok {
			return ok
		}
	}

	return true
}

func Part2(reports [][]uint64) int {
	total := 0
	for _, report := range reports {
		if reportIsSafeWithTolerance(report) {
			total += 1
		}
	}
	return total
}

func reportIsSafeWithTolerance(report []uint64) bool {
	for i := 0; i < len(report); i++ {
		ok := reportIsSafe(report)
		if ok {
			return true
		} else {
			ok = reportIsSafe(remove(report, i))
			if ok {
				return true
			}
		}
	}
	return false
}

func remove(s []uint64, index int) []uint64 {
	ret := make([]uint64, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
