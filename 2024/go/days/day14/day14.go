package day14

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

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
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	width := 101
	height := 103

	robots, err := LoadInput(s)
	if err != nil {
		return 0, err
	}

	return SafetyScore(robots, width, height), nil
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
	PX, PY, VX, VY int
}

// func PrintMap(m *common.TwoDMap[[]Robot]) {
// 	fmt.Print("\033[H\033[2J")
//
// 	for i, c := range m.Map {
// 		if len(c) > 0 {
// 			fmt.Printf("%d ", len(c))
// 		} else {
// 			fmt.Print(". ")
// 		}
// 		x := (i + 1) % int(m.Width)
// 		if x == 0 {
// 			fmt.Printf("\n")
// 		}
// 	}
// 	fmt.Printf("\n")
// }

func LoadInput(s *bufio.Scanner) ([]Robot, error) {
	robots := []Robot{}
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

		robots = append(robots, Robot{
			px, py, vx, vy,
		})
	}

	return robots, nil
}

func MoveRobots(rs []Robot, width, height int) {
	for i := 0; i < len(rs); i++ {
		rs[i].MoveRobot(width, height)
	}
}

func PrintRobots(robots []Robot, width, height int) {
	result := ""
	for y := range height {
		for x := range width {
			total := 0
			for _, r := range robots {
				if r.PX == x && r.PY == y {
					total += 1
				}
			}
			if total > 0 {
				result += fmt.Sprintf("%d ", total)
			} else {
				result += ". "
			}
			newLine := (x + 1) % width
			if newLine == 0 {
				result += "\n"
			}
		}
	}
	fmt.Printf("%v\n", result)
}

func (r *Robot) MoveRobot(width, height int) {
	r.PX += r.VX
	if r.PX >= width {
		r.PX -= width
	} else if r.PX < 0 {
		r.PX += width
	}

	r.PY += r.VY
	if r.PY >= height {
		r.PY -= height
	} else if r.PY < 0 {
		r.PY += height
	}
}

func SafetyScore(robots []Robot, width, height int) int {
	mw := FindMiddle(width)
	mh := FindMiddle(height)
	quadrants := map[int]int{1: 0, 2: 0, 3: 0, 4: 0}
	for _, r := range robots {
		switch {
		case r.PX < mw && r.PY < mh:
			quadrants[1] += 1
		case r.PX > mw && r.PY < mh:
			quadrants[2] += 1
		case r.PX < mw && r.PY > mh:
			quadrants[3] += 1
		case r.PX > mw && r.PY > mh:
			quadrants[4] += 1
		}
	}

	return quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]
}

func FindMiddle(length int) int {
	return (length - 1) / 2
}
