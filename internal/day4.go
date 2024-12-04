package internal

import (
	"fmt"
	"strings"
)

type Day4 struct{}

func countWords(input string) int {
	xmasCount := strings.Count(input, "XMAS")
	smaxCount := strings.Count(input, "SAMX")
	return xmasCount + smaxCount
}

func solvePartOne(words [][]string) int {
	c := 0
	for i := 0; i < len(words); i++ {
		straightLine := ""
		diagonalPass := ""
		mirrorDiagonal := ""
		reverseDiagonalPass := ""
		reverseMirrorDiagonal := ""
		verticalLine := ""
		for j := 0; j < len(words); j++ {
			straightLine += words[i][j]
			verticalLine += words[j][i]
			if i+j < len(words) {
				diagonalPass += words[i+j][j]
				if i > 0 {
					mirrorDiagonal += words[j][i+j]
					reverseMirrorDiagonal += words[j][len(words)-(i+j)-1]
				}
				reverseDiagonalPass += words[i+j][len(words)-1-j]
			}
		}
		c += countWords(straightLine)
		c += countWords(verticalLine)
		c += countWords(diagonalPass)
		c += countWords(mirrorDiagonal)
		c += countWords(reverseDiagonalPass)
		c += countWords(reverseMirrorDiagonal)
	}
	return c
}

func checkXMAS(x, y int, words [][]string) bool {
	maxBound := len(words) - 1
	if x-1 < 0 || y-1 < 0 || x+1 > maxBound || y+1 > maxBound {
		return false
	}
	diag1 := words[x-1][y-1] + words[x][y] + words[x+1][y+1]
	diag2 := words[x-1][y+1] + words[x][y] + words[x+1][y-1]
	return (diag1 == "SAM" || diag1 == "MAS") && (diag2 == "MAS" || diag2 == "SAM")
}

// Put dots in place of irrelevant characters
func solvePartTwo(words [][]string) int {
	c := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			center := words[i][j]
			if center != "A" {
				continue
			}
			if checkXMAS(i, j, words) {
				c++
			}
		}
	}
	return c
}

func (d Day4) Run(filepath string) {
	words := [][]string{}
	count := 0
	StreamFile(filepath, func(line string) {
		words = append(words, strings.Split(line, ""))
	})
	fmt.Printf("Shape: %dx%d\n", len(words), len(words[0]))
	count += solvePartOne(words)
	fmt.Println("Part 1:", count)

	// Part 2
	count = 0
	count += solvePartTwo(words)
	fmt.Println("Part 2:", count)
}
