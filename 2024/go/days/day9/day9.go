package day9

import (
	"bufio"
	// "fmt"
	"reflect"
	"strconv"

	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 9

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

	blocks := LoadInput(s)

	return Compress(blocks), nil
}

func Part2() (int, error) {
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	blocks := LoadInput(s)

	return CompressPartTwo(blocks), nil
}

func LoadInput(s *bufio.Scanner) []int {
	s.Scan()
	line := s.Text()

	blocks := []int{}
	isFile := true
	idIndex := 0

	for _, c := range line {
		v, _ := strconv.Atoi(string(c))
		id := -1
		if isFile {
			id = idIndex
			idIndex += 1
		}
		for range v {
			blocks = append(blocks, id)
		}
		isFile = !isFile
	}

	return blocks
}

func Compress(blocks []int) int {
	nextBlock := len(blocks) - 1
	checksum := 0
	for i := 0; i < len(blocks); i++ {
		curr := blocks[i]
		b := blocks[nextBlock]
		if b == -1 {
			nextBlock -= 1
			i -= 1
			continue
		} else if nextBlock < i {
			break
		}

		if curr == -1 {
			reflect.Swapper(blocks)(i, nextBlock)
		}

		curr = blocks[i]
		if curr > 0 {
			checksum += curr * i
		}
	}

	return checksum
}

type Block struct {
	Id     int
	Length int
}

func LoadInputPartTwo(s *bufio.Scanner) []Block {
	s.Scan()
	line := s.Text()

	blocks := []Block{}
	isFile := true
	idIndex := 0

	for _, c := range line {
		v, _ := strconv.Atoi(string(c))
		id := -1
		if isFile {
			id = idIndex
			idIndex += 1
		}
		if v > 0 {
			blocks = append(blocks, Block{Id: id, Length: v})
		}
		isFile = !isFile
	}

	return blocks
}

func CompressPartTwo(blocks []int) int {

	// fmt.Printf("%v\n", blocks)
	// seen := map[int]bool{}
	// Loop:
	for i := len(blocks) - 1; i >= 0; i-- {

		emptySpaceStart, _ := FindEmptySpace(0, blocks)
		block := blocks[i]
		if block == -1 {
			continue
		} else if i < emptySpaceStart {
			break
		}

		blockLength := FindChunk(i, block, blocks)

		for x := 0; x < i-blockLength; x++ {
			emptySpaceStart, emptySpaceLength := FindEmptySpace(x, blocks)
			if emptySpaceLength == 0 {
				break
			}
			if emptySpaceLength >= blockLength && emptySpaceStart < i {
				// _, ok := seen[block]
				// if ok {
				// 	break Loop
				// }
				for j := 0; j < blockLength; j++ {
					blocks[emptySpaceStart+j] = block
					blocks[i-j] = -1
				}
				// seen[block] = true
				break
			}

			x = (emptySpaceStart + emptySpaceLength) - 1
		}

		// fmt.Printf("Index: %d, Result: %v\n", i, blocks)
		// KEEP AT END
		result := i - blockLength
		i = result + 1
	}

	checksum := 0

	for i, b := range blocks {
		if b == -1 {
			continue
		}

		checksum += (b * i)
	}

	return checksum
}

func FindChunk(start, block int, blocks []int) int {
	length := 0
	for i := len(blocks) - 1; i >= 0; i-- {
		if i > start {
			continue
		}
		next := blocks[i]
		if next != block {
			break
		}
		length += 1
	}
	return length
}

func FindEmptySpace(from int, blocks []int) (int, int) {
	length := 0

	prevIndex := 0
	start := 0
	for i := from; i < len(blocks)-1; i++ {
		if prevIndex != 0 && i-prevIndex > 1 {
			break
		}

		block := blocks[i]
		if block != -1 {
			continue
		}

		if start == 0 {
			start = i
		}

		prevIndex = i
		length += 1
	}

	return start, length
}
