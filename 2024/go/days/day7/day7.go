package day7

import (
	"bufio"
	"fmt"
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
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	equations, err := LoadInput(s)
	if err != nil {
		return 0, err
	}

	return GetPart1Total(equations), nil
}

func GetPart1Total(equations []Equation) int {
	total := 0
	for _, equation := range equations {
		combinations := Combinations[len(equation.Values)-1]
		for _, c := range combinations {
			ok := equation.IsValid(c)
			if ok {
				total += equation.Result
				break
			}
		}
	}

	return total
}

func Part2() (int, error) {
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		panic(err)
	}
	defer closeFile()

	equations, err := LoadInput(s)
	if err != nil {
		return 0, err
	}
	res, err := GetPart2Total(equations)
	return res, err
}

func GetPart2Total(equations []Equation) (int, error) {
	total := 0
	for _, equation := range equations {
		combinations := CombinationsPartTwo[len(equation.Values)-1]
		for _, c := range combinations {
			ok, err := equation.IsValidPartTwo(c)
			if err != nil {
				return total, err
			}
			if ok {
				total += equation.Result
				break
			}
		}
	}

	return total, nil
}

const (
	Add = iota
	Multiply
	Combine
)

var Combinations map[int][][]int = map[int][][]int{}
var CombinationsPartTwo map[int][][]int = map[int][][]int{}

type Equation struct {
	Result int
	Values []int
}

func (e Equation) IsValid(combinations []int) bool {
	total := GetSum(combinations, e.Values)
	return total == e.Result
}

func GetSum(combinations []int, values []int) int {
	total := 0

	if len(values) == 1 {
		return values[0]
	}

	if len(combinations) == 0 {
		total, _ = strconv.Atoi(strconv.Itoa(values[0]) + strconv.Itoa(values[1]))
		return total
	}

	for i := 0; i < len(values); i++ {
		value := values[i]
		if i == 0 {
			total += value
			continue
		}

		operator := combinations[i-1]
		if operator == Add {
			total += value
		} else {
			total *= value
		}
	}
	return total
}

func (e Equation) IsValidPartTwo(combinations []int) (bool, error) {
	combineIndexes := []int{}
	for i, c := range combinations {
		if c == Combine {
			combineIndexes = append(combineIndexes, i)
		}
	}

	parts := [][]int{}
	combinationParts := [][]int{}
	remaining := 0
	for i := 0; i < len(combineIndexes); i++ {
		curr := combineIndexes[i] + 1
		slice := e.Values[remaining:curr]
		parts = append(parts, slice)
		combinationParts = append(combinationParts, combinations[remaining:curr-1])
		remaining += len(slice)
	}

	parts = append(parts, e.Values[remaining:])
	combinationParts = append(combinationParts, combinations[remaining:])

	fmt.Printf("Operators: %v\n", combinations)
	fmt.Printf("Combination Indexes: %v\n", combineIndexes)
	fmt.Printf("Values: %v\n", e.Values)
	fmt.Printf("Parts: %v\n", parts)
	fmt.Printf("Combination Parts: %v\n", combinationParts)
	fmt.Printf("\n")

	var result int
	var err error
	for i, p := range parts {
		sum := GetSum(combinationParts[i], p)
		result, err = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(sum))
		if err != nil {
			return false, err
		}
	}

	// var result int
	// var err error
	// for _, ci := range combineIndexes {
	// 	var split []int
	// 	if len(e.Values) == 2 {
	// 		split = e.Values
	// 	} else {
	// 		split = e.Values[:ci]
	// 	}
	// 	// fmt.Printf("%v\n", split)
	// 	sum := GetSum(combinations[:ci], split)
	// 	result, err = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(sum))
	// 	if err != nil {
	// 		return false, err
	// 	}
	// }
	return result == e.Result, nil
}

func GenerateCombinations(count int, operators []int) [][]int {
	combinations := [][]int{}

	if count == 0 {
		return combinations
	}

	var helper func(current []int)
	helper = func(current []int) {
		if len(current) == count {
			temp := make([]int, len(current))
			copy(temp, current)
			combinations = append(combinations, temp)
			return
		}

		for _, c := range operators {
			current = append(current, c)

			helper(current)

			current = current[:len(current)-1]
		}
	}

	helper([]int{})

	return combinations
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

		// Dodgy side effect :(
		possibleCombinations := len(values) - 1
		_, ok := Combinations[possibleCombinations]
		if !ok {
			Combinations[possibleCombinations] = GenerateCombinations(possibleCombinations, []int{Add, Multiply})

		}

		_, ok = CombinationsPartTwo[possibleCombinations]
		if !ok {
			CombinationsPartTwo[possibleCombinations] = GenerateCombinations(possibleCombinations, []int{Add, Multiply, Combine})
		}

		equations = append(equations, Equation{
			Result: result,
			Values: values,
		})
	}

	return equations, nil
}
