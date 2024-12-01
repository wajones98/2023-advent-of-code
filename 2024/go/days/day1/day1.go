package day1

import (
	"github.com/wajones98/advent-of-code/input"
)

func Run(day int) (uint, error) {
	s, closeFile, err := input.GetInput(day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	for s.Scan() {
		println(s.Text())
	}

	return 0, nil
}
