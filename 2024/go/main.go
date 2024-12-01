package main

import (
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
		println("HELLO DAY 1")
	}
}
