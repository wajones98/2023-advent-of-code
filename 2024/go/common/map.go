package common

import "fmt"

type TwoDMap struct {
	Map    []string
	Width  int
	Height int
}

func NewTwoDMap(width, height int) *TwoDMap {
	return &TwoDMap{
		Map:    make([]string, width*height),
		Width:  width,
		Height: height,
	}
}

func (m *TwoDMap) Put(x, y int, r string) error {
	err := m.checkBounds(x, y)
	if err != nil {
		return err
	}
	m.Map[m.getIndex(x, y)] = r
	return nil
}

func (m *TwoDMap) Get(x, y int) (string, error) {
	err := m.checkBounds(x, y)
	if err != nil {
		return "", err
	}
	return m.Map[m.getIndex(x, y)], nil
}

func (m *TwoDMap) getIndex(x, y int) int {
	return y*m.Width + x
}

func (m *TwoDMap) checkBounds(x, y int) error {
	if x > m.Width {
		return fmt.Errorf("%d is out of bounds %d", x, m.Width)
	} else if y > m.Height {
		return fmt.Errorf("%d is out of bounds %d", y, m.Height)
	}

	return nil
}

func (m *TwoDMap) FindPosition(i int) (int, int) {
	y := i / m.Width
	x := i % m.Width
	return x, y
}

func (m *TwoDMap) String() string {
	result := ""
	for i, c := range m.Map {
		result += c
		x := (i + 1) % int(m.Width)
		if x == 0 {
			result += "\n"
		}
	}
	result += "\n"
	return result
}
