package day5

import (
	"bufio"
	"fmt"
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
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		panic(err)
	}
	defer closeFile()

	rules, updates, err := LoadInput(s)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, update := range updates {
		value, ok := UpdateIsOkay(rules, update)
		if ok {
			total += value
		}
	}

	return total
}

func Part2() int {
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		panic(err)
	}
	defer closeFile()

	rules, updates, err := LoadInput(s)
	if err != nil {
		panic(err)
	}

	badUpdates := [][]int{}
	for _, update := range updates {
		_, ok := UpdateIsOkay(rules, update)
		if !ok {
			badUpdates = append(badUpdates, update)
		}
	}
	total := 0
	for _, update := range badUpdates {
		value, ok := UpdateIsOkay(rules, FixUpdate(rules, update))
		if !ok {
			panic("Report is not okay when it should be")
		}
		total += value
	}

	return total
}

func FixUpdate(rules map[int][]int, updates []int) []int {
	fixed := make([]int, len(updates))
	copy(fixed, updates)

	for {
		_, ok := UpdateIsOkay(rules, fixed)
		if ok {
			break
		}
		fmt.Printf("%v\n", fixed)
		for i, update := range fixed {
			rule := rules[update]
			subset := updates[:i]
			for _, r := range rule {
				for si, s := range subset {
					if s == r {
						fixed[i] = r
						fixed[si] = update
					}
				}
			}
		}
	}

	return fixed
}
