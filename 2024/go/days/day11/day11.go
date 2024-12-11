package day11

import (
	"bufio"
	"slices"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 11

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

func LoadInput(s *bufio.Scanner) ([]int, error) {
	s.Scan()
	parts := strings.Split(s.Text(), " ")
	values := []int{}
	for _, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func Blink(stones []int) []int {
	for i := 0; i < len(stones); i++ {
		stone := stones[i]
		transformed := TransformStone(stone)
		stones[i] = transformed[0]
		if len(transformed) > 1 {
			i += 1
			stones = slices.Insert(stones, i, transformed[1])
		}
	}

	return stones
}

func TransformStone(stone int) []int {
	stones := []int{}
	length := len(strconv.Itoa(stone))
	switch {
	case stone == 0:
		stones = append(stones, 1)
	case length%2 == 0:
		stoneString := strconv.Itoa(stone)
		first, _ := strconv.Atoi(stoneString[:length/2])
		second, _ := strconv.Atoi(stoneString[length/2:])
		stones = append(stones, first, second)
	default:
		stones = append(stones, stone*2024)
	}

	return stones
}
