package day9

import (
	"bufio"
	"fmt"
	"reflect"
	"slices"
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
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	return 0, nil
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

func CompressPartTwo(blocks []Block) int {
	blockIndex := len(blocks) - 1
	emptyIndex := 0

	checksum := 0

	fmt.Printf("%v\n", blocks)

	for {
		if emptyIndex == len(blocks) || blockIndex < emptyIndex {
			break
		}

		currBlock := Block{
			Id:     blocks[emptyIndex].Id,
			Length: blocks[emptyIndex].Length,
		}
		if currBlock.Id != -1 {
			emptyIndex += 1
			continue
		}

		blockToMove := blocks[blockIndex]

		if blockToMove.Length == currBlock.Length {
			reflect.Swapper(blocks)(blockIndex, emptyIndex)
		} else if blockToMove.Length < currBlock.Length {
			newBlock := Block{Id: -1, Length: blockToMove.Length}
			currBlock.Length -= blockToMove.Length
			blocks[emptyIndex] = blockToMove
			blocks[blockIndex] = newBlock
			blocks = slices.Insert(blocks, emptyIndex+1, currBlock)
			break
		}

	}

	fmt.Printf("%v\n", blocks)

	return checksum
}
