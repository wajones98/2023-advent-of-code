package day1

import (
	"slices"
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

	leftList := []int{}
	rightList := []int{}

	for s.Scan() {
		line := s.Text()
		values := strings.Split(line, "   ")
		left, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, err
		}
		right, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, err
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return &days.Result[int, int]{
		Part1: Part1(leftList, rightList),
		Part2: Part2(leftList, rightList),
	}, nil
}

func Part1(leftList, rightList []int) int {
	slices.Sort(leftList)
	slices.Sort(rightList)

	sum := 0
	for i := 0; i < len(leftList); i++ {
		left := leftList[i]
		right := rightList[i]

		if left > right {
			sum += (left - right)
		} else {
			sum += (right - left)
		}
	}

	return sum
}

func Part2(leftList, rightList []int) int {
	sum := 0
	for _, left := range leftList {
		frequency := 0
		for _, right := range rightList {
			if right == left {
				frequency += 1
			}
		}
		sum += (left * frequency)
	}
	return sum
}
