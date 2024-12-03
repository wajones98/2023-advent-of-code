package day3

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/wajones98/advent-of-code/input"

	"github.com/wajones98/advent-of-code/days"
)

const Day int = 3

type Instruction struct {
	Left, Right int
}

func Run() (*days.Result[int, int], error) {

	return &days.Result[int, int]{
		Part1: Part1(),
		Part2: 0,
	}, nil
}

func LoadLines() ([]string, error) {
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return nil, err
	}
	defer closeFile()
	lines := []string{}
	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func Part1() int {
	lines, err := LoadLines()
	if err != nil {
		panic(err)
	}
	total := 0
	for _, line := range lines {
		instructions, err := GetInstructions(line)
		if err != nil {
			panic(err)
		}
		total += GetSum(instructions)
	}
	return total
}

func GetInstructions(line string) ([]*Instruction, error) {
	instructions := []*Instruction{}
	exp, err := regexp.Compile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	if err != nil {
		return nil, err
	}

	m := exp.FindAll([]byte(line), -1)
	if m == nil {
		return nil, err
	}

	for _, m := range m {
		instruction, err := GetInstruction(m)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func GetInstruction(instruction []byte) (*Instruction, error) {
	exp, err := regexp.Compile(`[0-9]{1,3}`)
	if err != nil {
		return nil, err
	}
	n := exp.FindAll(instruction, -1)
	if n == nil || len(n) != 2 {
		return nil, errors.New("Invalid numbers")
	}
	left, err := strconv.Atoi(string(n[0]))
	if err != nil {
		return nil, err
	}
	right, err := strconv.Atoi(string(n[1]))
	if err != nil {
		return nil, err
	}
	return &Instruction{
		Left:  left,
		Right: right,
	}, nil
}

func GetSum(instructions []*Instruction) int {
	sum := 0
	for _, i := range instructions {
		sum += (i.Left * i.Right)
	}
	return sum
}
