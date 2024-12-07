package day7

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/input"

	"github.com/wajones98/advent-of-code/days"
)

const Day int = 7

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
		panic(err)
	}
	defer closeFile()

	return 0, nil
}

const (
	Add = iota
	Multiply
)

var Combinations map[int][]int = map[int][]int{}

type Equation struct {
	Result int
	Values []int
}

func LoadInput(s *bufio.Scanner) ([]Equation, error) {
	equations := []Equation{}
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, ":")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		values := []int{}
		valuesString := strings.Split(parts[1], " ")
		for _, v := range valuesString {
			if v == "" {
				continue
			}
			value, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			values = append(values, value)
		}

		equations = append(equations, Equation{
			Result: result,
			Values: values,
		})
	}

	return equations, nil
}
