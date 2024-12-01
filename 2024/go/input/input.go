package input

import (
	"bufio"
	"fmt"
	"os"
)

const BaseUrl string = "https://adventofcode.com/2024/day/%d/input"

func GetInput(day int) (*bufio.Reader, error) {
	filename := fmt.Sprintf("input/day%d.txt", day)

	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	return bufio.NewReader(fi), nil
}
