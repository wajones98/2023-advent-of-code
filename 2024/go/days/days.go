package days

import (
	"fmt"
	"os"
)

type Result[T any, V any] struct {
	Part1 T
	Part2 V
}

const FileTemplate = `package day%d

import (
	"bufio"

	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = %d 

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
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	return 0, nil
}

func Part2() (int, error) {
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	return 0, nil
}

func LoadInput(s *bufio.Scanner) (_, error) {
}
`

const TestFileTemplate = `package day%d

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const Input = ""
const Data = ""

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(Data, *actual) {
		t.Errorf("Expected: %%v, Actual: %%v\\n", Data, actual)
	}
}
`

func ScaffoldDay(day int) error {
	dayName := fmt.Sprintf("day%d", day)

	err := os.Mkdir("days/"+dayName, os.ModePerm)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("days/%s/%s.go", dayName, dayName)
	err = os.WriteFile(fileName, []byte(fmt.Sprintf(FileTemplate, day, day)), 0644)
	if err != nil {
		return err
	}

	testFileName := fmt.Sprintf("days/%s/%s_test.go", dayName, dayName)
	err = os.WriteFile(testFileName, []byte(fmt.Sprintf(TestFileTemplate, day)), 0644)
	if err != nil {
		return err
	}

	return nil
}
