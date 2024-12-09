package day9

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const Input = `2333133121414131402`

func TestLoadInput(t *testing.T) {
	input := "12345"
	s := bufio.NewScanner(strings.NewReader(input))

	actual := LoadInput(s)
	expected := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Got: %v\n", expected, actual)
	}
}

func TestCompress(t *testing.T) {
	data := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	expected := []int{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}
	expectedChecksum := 60
	checksum := Compress(data)

	if !reflect.DeepEqual(expected, data) {
		t.Errorf("Expected: %v, Got: %v\n", expected, data)
	}

	if expectedChecksum != checksum {
		t.Errorf("Expected: %d, Got: %d\n", expectedChecksum, checksum)
	}

}

func TestExampleInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	data := LoadInput(s)
	expected := 1928
	checksum := Compress(data)

	if expected != checksum {
		t.Errorf("Expected: %d, Got: %d\n", expected, checksum)
	}
}
