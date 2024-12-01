package input

import (
	"bufio"
	"fmt"
	"os"
)

const BaseUrl string = "https://adventofcode.com/2024/day/%d/input"

func GetInput(day int) (*bufio.Scanner, func() error, error) {
	filename := fmt.Sprintf("input/day%d.txt", day)

	fi, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	return bufio.NewScanner(fi), fi.Close, nil
}
