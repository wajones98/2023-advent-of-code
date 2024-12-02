package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/input"

	"github.com/wajones98/advent-of-code/days"
)

const Day int = 2

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

func Part2() int {
	return 0
}
