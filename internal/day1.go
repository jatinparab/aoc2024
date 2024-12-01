package internal

import (
	"fmt"
	"math"
	"sort"
)

type Day1 struct{}

func (d Day1) Run(filepath string) {
	list1 := []int{}
	list2 := []int{}
	ans := 0
	StreamFileColumnsInt(filepath, func(num1 int, num2 int) {
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	})
	sort.Ints(list1)
	sort.Ints(list2)
	for i, num := range list1 {
		ans += int(math.Abs(float64(num - list2[i])))
	}
	fmt.Println("Part 1:", ans)
	ans = 0
	frequencies := FrequencyMap(list2)
	for _, num := range list1 {
		ans += num * frequencies[num]
	}
	fmt.Println("Part 2:", ans)
}
