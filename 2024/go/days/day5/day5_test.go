package day5

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const Input string = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	_, _, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

}

type Expected struct {
	Ok    bool
	Value int
}

func TestUpdateIsOkay(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	rules, updates, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}
	expected := []Expected{{true, 61}, {true, 53}, {true, 29}, {false, 0}, {false, 0}, {false, 0}}
	total := 0
	for i, update := range updates {
		value, ok := UpdateIsOkay(rules, update)
		if ok != expected[i].Ok {
			t.Errorf("\n-----------\n%v\nGot: %t\nExpected: %t\n-----------\n", update, ok, expected[i].Ok)
		} else if value != expected[i].Value {
			t.Errorf("\n-----------\n%v\nGot: %d\nExpected: %d\n-----------\n", update, value, expected[i].Value)
		}
		if ok {
			total += value
		}
	}

	expectedTotal := 143
	if total != expectedTotal {

		t.Errorf("Got: %d\nExpected: %d\n", total, expectedTotal)
	}
}

func TestFixUpdate(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	rules, updates, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	// Take just the last 3 failing tests
	updates = updates[3:]

	expected := [][]int{{97, 75, 47, 61, 53}, {61, 29, 13}, {97, 75, 47, 29, 13}}
	for i, update := range updates {
		_, ok := UpdateIsOkay(rules, update)
		if !ok {
			fixedUpdate := FixUpdate(rules, update)
			if !reflect.DeepEqual(fixedUpdate, expected[i]) {
				t.Errorf("\n-----------\n%v\nGot: %v\nExpected: %v\n-----------\n", updates[i], fixedUpdate, expected[i])
			}
		}
	}
}
