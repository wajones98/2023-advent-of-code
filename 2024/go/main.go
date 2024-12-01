package main

import (
	"fmt"
	"github.com/wajones98/advent-of-code/days/day1"
	"log"
	"os"
	"strconv"
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
	}
}
