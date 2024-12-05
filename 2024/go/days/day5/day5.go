package day5

import (
	"bufio"
	"slices"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/input"

	"github.com/wajones98/advent-of-code/days"
)

const Day int = 5

func Run() (*days.Result[int, int], error) {

	return &days.Result[int, int]{
		Part1: Part1(),
		Part2: Part2(),
	}, nil
}

func LoadInput(s *bufio.Scanner) (map[int][]int, [][]int, error) {
	rules := map[int][]int{}
	updates := [][]int{}
	isSectionOne := true
	for s.Scan() {
		line := s.Text()
		if line == "" || line == " " {
			isSectionOne = false
			continue
		}

		if isSectionOne {
			parts := strings.Split(line, "|")
			key, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, err
			}

			value, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}

			values, ok := rules[key]
			if !ok {
				values = []int{}
			}
			rules[key] = append(values, value)
		} else {
			parts := strings.Split(line, ",")
			values := []int{}
			for _, part := range parts {
				value, err := strconv.Atoi(part)
				if err != nil {
					return nil, nil, err
				}
				values = append(values, value)
			}
			updates = append(updates, values)
		}
	}
	return rules, updates, nil
}

func UpdateIsOkay(rules map[int][]int, updates []int) (int, bool) {
	for i, update := range updates {
		rule := rules[update]
		subset := updates[:i]
		for _, r := range rule {
			if slices.Contains(subset, r) {
				return 0, false
			}
		}
	}
	return updates[len(updates)/2], true
}

func Part1() int {
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		panic(err)
	}
	defer closeFile()

	return 0
}

func Part2() int {
	return 0
}
