package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/wajones98/advent-of-code/days/day1"
	"github.com/wajones98/advent-of-code/days/day10"
	"github.com/wajones98/advent-of-code/days/day11"
	"github.com/wajones98/advent-of-code/days/day2"
	"github.com/wajones98/advent-of-code/days/day3"
	"github.com/wajones98/advent-of-code/days/day4"
	"github.com/wajones98/advent-of-code/days/day5"
	"github.com/wajones98/advent-of-code/days/day6"
	"github.com/wajones98/advent-of-code/days/day7"
	"github.com/wajones98/advent-of-code/days/day8"
	"github.com/wajones98/advent-of-code/days/day9"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("Expected 1 argument. Got %d args", len(args))
	}

	arg, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Expected argument to be integer, got %s", args[0])
	}

	switch arg {
	case 1:
		result, err := day1.Run(arg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)
	case 2:
		result, err := day2.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)
	case 3:
		result, err := day3.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)
	case 4:
		result, err := day4.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)

	case 5:
		result, err := day5.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)

	case 6:
		result, err := day6.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)
	case 7:
		result, err := day7.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)
	case 8:
		result, err := day8.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)
	case 9:
		result, err := day9.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)
	case 10:
		result, err := day10.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)
	case 11:
		result, err := day11.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 Result: %d\n", result.Part1)
		fmt.Printf("Part 2 Result: %d\n", result.Part2)
	default:
		fmt.Printf("%d is not a valid day\n", arg)
	}
}
