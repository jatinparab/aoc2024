package internal

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day5 struct{}

func fixUpdateOrder(update []string, forwardMap map[string][]string) []string {
	for i := 0; i < len(update); i++ {
		for j := 0; j < len(update); j++ {
			if i == j {
				continue
			}
			first := update[i]
			second := update[j]
			if _, ok := forwardMap[first]; !ok {
				continue
			}
			for _, rule := range forwardMap[first] {
				if rule == second {
					update[i], update[j] = update[j], update[i]
				}
			}
		}
	}
	slices.Reverse(update)
	return update
}

func (d Day5) Run(filepath string) {
	rules := []string{}
	updates := [][]string{}
	forwardMap := map[string][]string{}
	appendUpdates := false
	StreamFile(filepath, func(line string) {
		if line == "" {
			appendUpdates = true
			return
		}
		if appendUpdates {
			updates = append(updates, strings.Split(line, ","))
			return
		}
		rules = append(rules, line)
	})
	for _, rule := range rules {
		split := strings.Split(rule, "|")
		forwardMap[split[1]] = append(forwardMap[split[1]], split[0])
	}
	invalidUpdateIdxs := map[int]bool{}
	for updateIdx, update := range updates {
		for i := 0; i < len(update); i++ {
			for j := i; j < len(update); j++ {
				first := update[i]
				second := update[j]
				if i == j {
					continue
				}
				if _, ok := forwardMap[first]; !ok {
					continue
				}
				for _, rule := range forwardMap[first] {
					if rule == second {
						invalidUpdateIdxs[updateIdx] = true
					}
				}
			}
		}
	}
	part1ans := 0
	for i, update := range updates {
		if _, ok := invalidUpdateIdxs[i]; ok {
			continue
		}
		middleElem := len(update) / 2
		middle, _ := strconv.Atoi(update[middleElem])
		part1ans += middle
	}
	fmt.Println("Part 1:", part1ans)

	// Part 2
	part2ans := 0
	for invalidUpdateIdx := range invalidUpdateIdxs {
		fixed := fixUpdateOrder(updates[invalidUpdateIdx], forwardMap)
		middle := len(fixed) / 2
		middleInt, _ := strconv.Atoi(fixed[middle])
		part2ans += middleInt
	}
	fmt.Println("Part 2:", part2ans)
}
