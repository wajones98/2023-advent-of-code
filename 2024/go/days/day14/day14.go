package day14

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/common"
	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 14

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

type Robot struct {
	X, Y int
}

func LoadInput(s *bufio.Scanner, width, height int) (*common.TwoDMap[[]Robot], error) {
	twoDMap := common.NewTwoDMap[[]Robot](width, height)

	for s.Scan() {
		pv := strings.Split(s.Text(), " ")

		p := strings.Split(strings.ReplaceAll(pv[0], "p=", ""), ",")
		px, err := strconv.Atoi(p[0])
		if err != nil {
			return nil, err
		}
		py, err := strconv.Atoi(p[1])
		if err != nil {
			return nil, err
		}

		v := strings.Split(strings.ReplaceAll(pv[1], "v=", ""), ",")
		vx, err := strconv.Atoi(v[0])
		if err != nil {
			return nil, err
		}

		vy, err := strconv.Atoi(v[1])
		if err != nil {
			return nil, err
		}

		r, err := twoDMap.Get(px, py)
		if err != nil {
			return nil, err
		}
		r = append(r, Robot{vx, vy})

		err = twoDMap.Put(px, py, r)
		if err != nil {
			return nil, err
		}
	}

	return twoDMap, nil
}
