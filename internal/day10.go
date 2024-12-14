package internal

import (
	"fmt"
	"strconv"
)

type Day10 struct {
	grid      [][]string
	peaks     map[string]bool
	sumRating int
}

func (d *Day10) Run(filepath string) {
	grid := [][]string{}
	StreamFile(filepath, func(line string) {
		row := []string{}
		for _, c := range line {
			row = append(row, string(c))
		}
		grid = append(grid, row)
	})
	d.grid = grid
	score := 0
	for i := 0; i < len(d.grid); i++ {
		for j := 0; j < len(d.grid[i]); j++ {
			if d.grid[i][j] == "0" {
				fscore := d.TrailScore(i, j)
				score += fscore
				if score > 0 {
					d.peaks = map[string]bool{}
				}
			}
		}
	}
	println("Part 1:", score)
	fmt.Println("Part 2:", d.sumRating)
}

func (d *Day10) isWithinGrid(x int, y int) bool {
	return x >= 0 && x < len(d.grid) && y >= 0 && y < len(d.grid[0])
}

func (d *Day10) isGradual(current string, next string) bool {
	if next == "" {
		return false
	}
	currentInt, _ := strconv.Atoi(current)
	nextInt, _ := strconv.Atoi(next)
	return nextInt-currentInt == 1
}

func (d *Day10) TrailScore(x int, y int) int {
	current := d.grid[x][y]
	if current == "9" {
		d.sumRating += 1
		key := fmt.Sprintf("%d,%d", x, y)
		if d.peaks[key] {
			return 0
		}
		d.peaks[key] = true
		return 1
	}
	choices := [][]int{}
	if d.isWithinGrid(x, y-1) {
		choices = append(choices, []int{x, y - 1})
	}
	if d.isWithinGrid(x, y+1) {
		choices = append(choices, []int{x, y + 1})
	}
	if d.isWithinGrid(x-1, y) {
		choices = append(choices, []int{x - 1, y})
	}
	if d.isWithinGrid(x+1, y) {
		choices = append(choices, []int{x + 1, y})
	}
	score := 0
	for _, choice := range choices {
		nextX, nextY := choice[0], choice[1]
		if d.isGradual(current, d.grid[nextX][nextY]) {
			nextScore := d.TrailScore(nextX, nextY)
			score += nextScore
		}
	}
	return score
}

func NewDay10() *Day10 {
	return &Day10{
		grid:      [][]string{},
		peaks:     map[string]bool{},
		sumRating: 0,
	}
}
