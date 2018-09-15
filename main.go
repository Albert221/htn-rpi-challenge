package main

import (
	"bufio"
	"fmt"
	"strings"
)

var input = `1 2 3
8 9 4
7 6 5`

func main() {
	parsed := parse(input)
	output := decode(parsed)

	for _, out := range output {
		fmt.Print(string(out) + " ")
	}
}

func parse(input string) [][]rune {
	var matrix [][]rune
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)

		var lineArr []rune
		for lineScanner.Scan() {
			word := lineScanner.Text()

			char := []rune(word)[0]
			lineArr = append(lineArr, char)
		}

		matrix = append(matrix, lineArr)
	}

	return matrix
}

const (
	Right = 0
	Down  = 1
	Left  = 2
	Up    = 3
)

func decode(input [][]rune) []rune {
	var result []rune
	width := len(input[0])
	height := len(input)

	direction := Right
	for {
		switch direction {
		case Right:
			result = append(result, input[0]...)
			// Removing top row
			input = input[1:]
			height--

		case Left:
			for i := width - 1; i >= 0; i-- {
				result = append(result, input[height-1][i])
			}
			// Remove bottom row
			input = input[:height-1]
			height--

		case Down:
			for _, line := range input {
				result = append(result, line[width-1])
			}
			// Remove right column
			for i, line := range input {
				input[i] = line[:width-1]
			}
			width--

		case Up:
			for i := height - 1; i >= 0; i-- {
				result = append(result, input[0][i])
			}
			// Remove left column
			for i, line := range input {
				input[i] = line[1:]
			}
			width--
		}

		direction = (direction + 1) % 4

		if width == 0 || height == 0 {
			break
		}
	}

	return result
}
