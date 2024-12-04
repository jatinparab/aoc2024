package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type Day3 struct {
	mulEnabled bool
}

func verifyIntSize(num int) bool {
	return num >= 0 && num <= 999
}

func verifyAndMultiply(input string) (int, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid input: %s", input)
	}
	num1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	num2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	if !verifyIntSize(num1) || !verifyIntSize(num2) {
		return 0, fmt.Errorf("invalid number: %d,%d", num1, num2)
	}
	return num1 * num2, nil
}

func (d *Day3) ParseAndToggleMul(part string) {
	for i := 0; i < len(part); i++ {
		if strings.HasPrefix(part[i:], "do()") {
			d.mulEnabled = true
		}
		if strings.HasPrefix(part[i:], "don't()") {
			d.mulEnabled = false
		}
	}
}

func (d *Day3) solve(input string, parseDoAndDonts bool) int {
	mulparts := strings.Split(input, "mul")
	part1Ans := 0
	for _, part := range mulparts {
		if part[0] != '(' {
			if parseDoAndDonts {
				d.ParseAndToggleMul(part)
			}
			continue
		}
		closingBracketIdx := strings.Index(part, ")")
		if closingBracketIdx == -1 {
			if parseDoAndDonts {
				d.ParseAndToggleMul(part)
			}
			continue
		}
		inside := part[1:closingBracketIdx]
		num, err := verifyAndMultiply(inside)
		if err != nil {
			if parseDoAndDonts {
				d.ParseAndToggleMul(part)
			}
			continue
		}
		if d.mulEnabled {
			part1Ans += num
		}
		// After mul operation
		if parseDoAndDonts {
			d.ParseAndToggleMul(part[closingBracketIdx+1:])
		}
	}
	return part1Ans
}

func (d *Day3) Run(filepath string) {
	part1Ans := 0
	StreamFile(filepath, func(line string) {
		part1Ans += d.solve(line, false)
	})
	fmt.Println("Part 1:", part1Ans)
	// Part 2
	part2Ans := 0
	StreamFile(filepath, func(line string) {
		part2Ans += d.solve(line, true)
	})
	fmt.Println("Part 2:", part2Ans)
}

func NewDay3() *Day3 {
	return &Day3{mulEnabled: true}
}
