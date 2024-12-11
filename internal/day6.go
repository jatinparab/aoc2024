package internal

import (
	"fmt"
	"strings"
	"time"
)

type Day6 struct {
	grid          [][]string
	playerPos     []int
	currDir       string
	visitedCount  int
	nTurns        int
	turningPoints [][]int
	state         []string
	repeating     int
}

func (d *Day6) play() error {
	debug := false
	for {
		err := d.move()
		if err != nil {
			break
		}
		if debug {
			fmt.Println()
			d.PrintBoard()
			fmt.Println()
			time.Sleep(50 * time.Millisecond)
			fmt.Print("\033[H\033[2J")
		}
		if d.checkStateForRepeats() {
			if d.repeating == 4 {
				fmt.Printf("Turns, %v\n", d.turningPoints)
				return fmt.Errorf("repeating")
			}
		}
	}
	return nil
}

func (d *Day6) reset() {
	d.currDir = "^"
	d.nTurns = 0
	d.turningPoints = [][]int{}
	d.visitedCount = 1
	d.grid = [][]string{}
	d.playerPos = []int{0, 0}
	d.state = []string{}
	d.repeating = 0
}

func (d *Day6) PrintBoard() {
	for i := 0; i < len(d.grid); i++ {
		for j := 0; j < len(d.grid[i]); j++ {
			fmt.Print(d.grid[i][j])
		}
		fmt.Println()
	}
}

// Checks if last 4 and the previous 4 to that are the same
func checkRectRepeat(turningPoints [][]int) bool {
	fmt.Println("Checking for rect repeat", turningPoints)
	if len(turningPoints) < 8 {
		return false
	}
	for i := 0; i < 4; i++ {
		if turningPoints[len(turningPoints)-1-i][0] == turningPoints[len(turningPoints)-5-i][0] && turningPoints[len(turningPoints)-1-i][1] == turningPoints[len(turningPoints)-5-i][1] {
			return true
		}
	}
	return false
}

func (d *Day6) Run(filepath string) {
	d.initGrid(filepath)
	err := d.play()

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Part 1:", d.visitedCount, d.nTurns)

	gridSizeX, gridSizeY := len(d.grid), len(d.grid[0])
	goodObstacles := 0

	// Part 2
	for i := 0; i < gridSizeX; i++ {
		for j := 0; j < gridSizeY; j++ {
			d.reset()
			d.initGrid(filepath)
			d.placeObstacle(i, j)
			err := d.play()
			if err != nil {
				fmt.Println("Error:", err)
				if strings.Contains(err.Error(), "repeating") {
					if checkRectRepeat(d.turningPoints) {
						goodObstacles++
					}
				}

			}
		}
	}
	fmt.Printf("Part 2: %d\n", goodObstacles)
}
func (d *Day6) printLastNTurningPoints(n int) {
	if len(d.turningPoints) < n {
		return
	}
	for i := 0; i < n; i++ {
		fmt.Println(d.turningPoints[len(d.turningPoints)-1-i])
	}
}

func (d *Day6) checkStateForRepeats() bool {
	for _, state := range d.state {
		if state == fmt.Sprintf("%s%d,%d", d.currDir, d.playerPos[0], d.playerPos[1]) {
			fmt.Printf("Repeating state: %s\n", state)
			return true
		}
	}
	return false
}

func checkTurnsForLoop(turningPoints [][]int) bool {
	if len(turningPoints) < 8 {
		return false
	}

	for i := 0; i < 4; i++ {
		if turningPoints[len(turningPoints)-1-i][0] != turningPoints[len(turningPoints)-5-i][0] || turningPoints[len(turningPoints)-1-i][1] != turningPoints[len(turningPoints)-5-i][1] {
			return false
		}
	}
	return true
}

func (d *Day6) placeObstacle(i, j int) {
	fmt.Printf("Placing obstacle at %d, %d\n", i, j)
	if d.grid[i][j] == "#" {
		return
	}
	d.grid[i][j] = "#"
}

func (d *Day6) initGrid(filepath string) {
	i := 0
	StreamFile(filepath, func(line string) {
		row := []string{}
		for j, char := range line {
			row = append(row, string(char))
			if string(char) == "^" {
				d.currDir = string(char)
				d.playerPos = []int{i, j}
			}
		}
		d.grid = append(d.grid, row)
		d.grid[d.playerPos[0]][d.playerPos[1]] = "X"
		i++
	})
}

func (d *Day6) isOnEdge() bool {
	return d.playerPos[0] == 0 || d.playerPos[0] == len(d.grid)-1 || d.playerPos[1] == 0 || d.playerPos[1] == len(d.grid[0])-1
}

func (d *Day6) isFacingObstacle() bool {
	switch d.currDir {
	case "^":
		return d.grid[d.playerPos[0]-1][d.playerPos[1]] == "#"
	case "<":
		return d.grid[d.playerPos[0]][d.playerPos[1]-1] == "#"
	case ">":
		return d.grid[d.playerPos[0]][d.playerPos[1]+1] == "#"
	case "V":
		return d.grid[d.playerPos[0]+1][d.playerPos[1]] == "#"
	}
	return false
}

func (d *Day6) markVisited() bool {
	if d.grid[d.playerPos[0]][d.playerPos[1]] != "X" {
		d.grid[d.playerPos[0]][d.playerPos[1]] = "X"
		d.visitedCount++
		return true
	}
	return false
}

func (d *Day6) turnRight() {
	switch d.currDir {
	case "^":
		d.currDir = ">"
	case "<":
		d.currDir = "^"
	case ">":
		d.currDir = "V"
	case "V":
		d.currDir = "<"
	}
	d.nTurns++
	playerPosCopy := make([]int, len(d.playerPos))
	copy(playerPosCopy, d.playerPos)
	d.turningPoints = append(d.turningPoints, playerPosCopy)
	if d.checkStateForRepeats() {
		d.repeating++
	}
}

func (d *Day6) movePlayer(direction string) error {
	switch direction {
	case "^":
		d.playerPos[0]--
	case "<":
		d.playerPos[1]--
	case ">":
		d.playerPos[1]++
	case "V":
		d.playerPos[0]++
	}
	if d.isOnEdge() {
		d.markVisited()
		return fmt.Errorf("player will go out of bounds")
	}
	return nil
}

func (d *Day6) move() error {
	if d.isFacingObstacle() {
		d.turnRight()
	}
	d.state = append(d.state, fmt.Sprintf("%s%d,%d", d.currDir, d.playerPos[0], d.playerPos[1]))
	d.markVisited()
	return d.movePlayer(d.currDir)
}

func NewDay6() *Day6 {
	return &Day6{
		grid:         [][]string{},
		playerPos:    []int{0, 0},
		currDir:      "^",
		visitedCount: 1,
	}
}
