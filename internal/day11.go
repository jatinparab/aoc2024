package internal

import (
	"fmt"
	"strconv"
)

type Day11 struct {
	memoTable map[string]int
}

func (d *Day11) calculateSteps(num int, nSteps int) int {
	new := processNumber(num)
	if nSteps == 1 {
		return len(new)
	}
	ans := 0
	for _, n := range new {
		key := fmt.Sprintf("%d-%d", n, nSteps-1)
		if _, ok := d.memoTable[key]; !ok {
			d.memoTable[key] = d.calculateSteps(n, nSteps-1)
		}
		ans += d.memoTable[key]
	}
	return ans
}

func processNumber(num int) []int {
	if num == 0 {
		return []int{1}
	}
	numStr := strconv.Itoa(num)
	if len(numStr)%2 == 0 {
		mid := len(numStr) / 2
		left, _ := strconv.Atoi(numStr[:mid])
		right, _ := strconv.Atoi(numStr[mid:])
		return []int{left, right}
	}
	return []int{num * 2024}
}

func (d *Day11) Run(filepath string) {

	StreamFileInts(filepath, func(nums []int) {
		ans := 0
		ans2 := 0
		for _, num := range nums {
			ans += d.calculateSteps(num, 25)
			ans2 += d.calculateSteps(num, 75)
		}
		println("Part 1:", ans)
		println("Part 2:", ans2)
	})

}

func NewDay11() *Day11 {
	return &Day11{
		memoTable: map[string]int{},
	}
}
