package common

import "fmt"

type TwoDMap[T any] struct {
	Map    []T
	Width  int
	Height int
}

func NewTwoDMap[T any](width, height int) *TwoDMap[T] {
	return &TwoDMap[T]{
		Map:    make([]T, width*height),
		Width:  width,
		Height: height,
	}
}

func (m *TwoDMap[T]) Put(x, y int, r T) error {
	err := m.CheckBounds(x, y)
	if err != nil {
		return err
	}
	m.Map[m.getIndex(x, y)] = r
	return nil
}

func (m *TwoDMap[T]) Get(x, y int) (T, error) {
	err := m.CheckBounds(x, y)
	var empty T
	if err != nil {
		return empty, err
	}
	return m.Map[m.getIndex(x, y)], nil
}

func (m *TwoDMap[T]) getIndex(x, y int) int {
	return y*m.Width + x
}

func (m *TwoDMap[T]) CheckBounds(x, y int) error {
	if x > m.Width || x < 0 {
		return fmt.Errorf("%d is out of bounds %d", x, m.Width)
	} else if y > m.Height || y < 0 {
		return fmt.Errorf("%d is out of bounds %d", y, m.Height)
	}

	return nil
}

func (m *TwoDMap[T]) FindPosition(i int) (int, int) {
	y := i / m.Width
	x := i % m.Width
	return x, y
}

func (m *TwoDMap[T]) String() string {
	result := ""
	for i, c := range m.Map {
		result += fmt.Sprintf("%v", c)
		x := (i + 1) % int(m.Width)
		if x == 0 {
			result += "\n"
		}
	}
	result += "\n"
	return result
}
