package day10

import (
	"bufio"
	"strings"
	"testing"
)

const Input = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestLoadInput(t *testing.T) {
	input := `0123
	1234
	8765
	9876`

	s := bufio.NewScanner(strings.NewReader(input))
	twoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}
	if strings.ReplaceAll(twoDMap.String(), "\n", "") != strings.ReplaceAll(Input, "\n", "") {
		t.Errorf("\nExpected: \n%s\nGot: \n%s\n", Input, twoDMap.String())
	}
}
