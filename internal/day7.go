package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type Day7 struct{}

func stringEqToList(s string) []string {
	eq := []string{}
	num := ""
	for _, c := range s {
		if c == '+' || c == '*' || c == '|' {
			eq = append(eq, num)
			eq = append(eq, string(c))
			num = ""
		} else {
			num += string(c)
		}
	}
	eq = append(eq, num)
	return eq
}

func calculate(equation string) int {
	parts := stringEqToList(equation)
	ans := 0
	op := "+"
	for _, part := range parts {
		if part == "+" || part == "*" || part == "|" {
			op = part
		} else {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Error converting %s to int\n", part)
			}
			if op == "+" {
				ans += num
			} else if op == "*" {
				ans *= num
			} else if op == "|" {
				stringConcat := fmt.Sprintf("%d%d", ans, num)
				ans, _ = strconv.Atoi(stringConcat)
			}
		}
	}
	return ans
}

func createEquations(numberLine string, isPart2 bool) []string {
	operators := []string{"+", "*"}
	if isPart2 {
		operators = append(operators, "|")
	}
	equations := []string{}
	numbers := strings.Split(numberLine, " ")
	if len(numbers) == 1 {
		return numbers
	}
	restEquations := createEquations(
		strings.Join(numbers[1:], " "),
		isPart2,
	)
	for _, operator := range operators {
		for _, restEquation := range restEquations {
			equations = append(equations, numbers[0]+operator+restEquation)
		}
	}
	return equations
}

func (d Day7) Run(file string) {
	p1Ans := 0
	p2Ans := 0
	StreamFile(
		file,
		func(s string) {
			parts := strings.Split(s, ":")
			answer, _ := strconv.Atoi(parts[0])
			numbers := parts[1]
			numbers = strings.TrimSpace(numbers)
			p1Equations := createEquations(numbers, false)
			p2Equations := createEquations(numbers, true)
			for _, equation := range p1Equations {
				if calculate(equation) == answer {
					p1Ans += answer
					break
				}
			}
			for _, equation := range p2Equations {
				if calculate(equation) == answer {
					p2Ans += answer
					break
				}
			}
		},
	)
	fmt.Println("Part 1:", p1Ans)
	fmt.Println("Part 2:", p2Ans)
}
