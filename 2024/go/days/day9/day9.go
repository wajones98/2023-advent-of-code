package day9

import (
	"bufio"
	"strconv"

	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 9

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

func LoadInput(s *bufio.Scanner) []int {
	s.Scan()
	line := s.Text()

	blocks := []int{}
	isFile := true
	idIndex := 0

	for _, c := range line {
		v, _ := strconv.Atoi(string(c))
		id := -1
		if isFile {
			id = idIndex
			idIndex += 1
		}
		for range v {
			blocks = append(blocks, id)
		}
		isFile = !isFile
	}

	return blocks
}

func Compress(blocks []int) {
	for i := 0; i < len(blocks); i++ {

	}
}
