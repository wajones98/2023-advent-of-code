package day5

import (
	"bufio"
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

func TestUpdateIsOkay(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	rules, updates, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}
	expected := []bool{true, true, true, false, false, false}
	for i, update := range updates {
		ok := UpdateIsOkay(rules, update)
		if ok != expected[i] {
			t.Errorf("-----------\n%v\nGot: %t\nExpected: %t\n-----------\n", update, ok, expected[i])
		}
	}
}
