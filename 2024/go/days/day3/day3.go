package day3

import (
	"regexp"
	"strconv"

	"github.com/wajones98/advent-of-code/input"

	"github.com/wajones98/advent-of-code/days"
)

const Day int = 3

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

type Instruction struct {
	Left, Right int
}

func GetInstructions(line string) []Instruction {
	instructions := []Instruction{}
	exp, err := regexp.Compile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	if err != nil {
		panic(err)
	}

	m := exp.FindAll([]byte(line), -1)
	if m == nil {
		return instructions
	}

	for _, m := range m {
		instructions = append(instructions, GetInstruction(m))
	}

	return instructions
}

func GetInstruction(instruction []byte) Instruction {
	exp, err := regexp.Compile(`[0-9]{1,3}`)
	if err != nil {
		panic(err)
	}
	n := exp.FindAll(instruction, -1)
	if n == nil || len(n) != 2 {
		panic("Invalid numbers")
	}
	left, err := strconv.Atoi(string(n[0]))
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(string(n[1]))
	if err != nil {
		panic(err)
	}
	return Instruction{
		Left:  left,
		Right: right,
	}
}

func Part1() int {
	return 0
}
