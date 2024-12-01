package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/input"
)

func Run(day int) (int, error) {
	s, closeFile, err := input.GetInput(day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	leftList := []int{}
	rightList := []int{}

	for s.Scan() {
		line := s.Text()
		values := strings.Split(line, "   ")
		fmt.Println(values)
		left, err := strconv.Atoi(values[0])
		if err != nil {
			return 0, err
		}
		right, err := strconv.Atoi(values[1])
		if err != nil {
			return 0, err
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
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

	return sum, nil
}
