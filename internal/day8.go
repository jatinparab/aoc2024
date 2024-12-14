package internal

import (
	"fmt"
	"math"
	"strings"
)

func CalculateDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

type Day8 struct {
	grid [][]string
}

type frequency string

type Antenna struct {
	positionX int
	positionY int
	frequency frequency
}

type Antinode struct {
	positionX int
	positionY int
}

type Point interface {
	GetPosition() (int, int)
}

func (a Antenna) GetPosition() (int, int) {
	return a.positionX, a.positionY
}

func (a Antinode) GetPosition() (int, int) {
	return a.positionX, a.positionY
}

func isInGrid(point Point, grid [][]string) bool {
	x, y := point.GetPosition()
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func (d *Day8) getPointsOnLine(antenna1, antenna2 Antenna, isPart1 bool) []Antinode {
	x1, y1 := antenna1.positionX, antenna1.positionY
	x2, y2 := antenna2.positionX, antenna2.positionY

	antinodes := []Antinode{}

	dx := x2 - x1
	dy := y2 - y1

	gcd := gcd(int(math.Abs(float64(dx))), int(math.Abs(float64(dy))))
	dx /= gcd
	dy /= gcd

	x, y := x1, y1
	for isInGrid(Antinode{
		positionX: x + dx,
		positionY: y + dy,
	}, d.grid) {
		x += dx
		y += dy
		if isPart1 && x == x2 && y == y2 {
			continue
		}
		antinodes = append(antinodes, Antinode{
			positionX: x,
			positionY: y,
		})
		if isPart1 {
			break
		}
	}
	x, y = x2, y2
	for isInGrid(Antinode{
		positionX: x - dx,
		positionY: y - dy,
	}, d.grid) {
		x -= dx
		y -= dy
		if isPart1 && x == x1 && y == y1 {
			continue
		}
		antinodes = append(antinodes, Antinode{
			positionX: x,
			positionY: y,
		})
		if isPart1 {
			break
		}
	}
	return antinodes
}

// func printGrid(sizeX int, sizeY int, antinodes []Antinode, antennas []Antenna) {
// 	grid := make([][]string, sizeX)
// 	for i := range grid {
// 		grid[i] = make([]string, sizeY)
// 		for j := range grid[i] {
// 			grid[i][j] = "."
// 		}
// 	}
// 	for _, antenna := range antennas {
// 		grid[antenna.positionX][antenna.positionY] = string(antenna.frequency)
// 	}
// 	for _, antinode := range antinodes {
// 		if isInGrid(antinode, grid) {
// 			grid[antinode.positionX][antinode.positionY] = "#"
// 		}
// 	}
// 	for _, row := range grid {
// 		fmt.Println(strings.Join(row, ""))
// 	}
// }

func (d *Day8) Run(file string) {
	grid := [][]string{}
	gridMap := map[frequency][]Antenna{}
	StreamFile(file, func(line string) {
		grid = append(grid, strings.Split(line, ""))
	})
	d.grid = grid
	for i, row := range grid {
		for j, cell := range row {
			if cell != "." {
				gridMap[frequency(cell)] = append(gridMap[frequency(cell)], Antenna{
					positionX: i,
					positionY: j,
					frequency: frequency(cell),
				})
			}
		}
	}
	p1Antinodes := []Antinode{}
	p2Antinodes := []Antinode{}
	for _, antennas := range gridMap {
		pairs := Combinations(antennas, 2)
		for _, pair := range pairs {
			p1Antinodes = append(p1Antinodes, d.getPointsOnLine(pair[0], pair[1], true)...)
			p2Antinodes = append(p2Antinodes, d.getPointsOnLine(pair[0], pair[1], false)...)
		}
	}
	fmt.Println("Part 1:", countUniqueAntinodes(p1Antinodes))
	fmt.Println("Part 2:", countUniqueAntinodes(p2Antinodes))
}

func countUniqueAntinodes(antinodes []Antinode) int {
	antinodeMap := map[Antinode]bool{}
	for _, antinode := range antinodes {
		antinodeMap[antinode] = true
	}
	return len(antinodeMap)
}

func NewDay8() *Day8 {
	return &Day8{
		grid: [][]string{},
	}
}
