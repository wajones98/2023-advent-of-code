package day3

import (
	"log"
	"testing"
)

func TestGetInstructions(t *testing.T) {
	data := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := []*Instruction{
		{
			Left:  2,
			Right: 4,
		},
		{
			Left:  5,
			Right: 5,
		},
		{
			Left:  11,
			Right: 8,
		},
		{
			Left:  8,
			Right: 5,
		},
	}
	actual, err := GetInstructions(data)
	if err != nil {
		panic(err)
	}

	if len(expected) != len(actual) {
		log.Fatalf("Expected Length: %v\nGot: %v\n", len(expected), len(actual))
	}

	for i, e := range expected {
		if *actual[i] != *e {
			log.Fatalf("Expected: %v\nGot: %v\n", expected, actual)
		}
	}
}
