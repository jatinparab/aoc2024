package internal

import (
	"fmt"
	"math"
)

type Day2 struct{}

func analyseLine(levels []int, dampen bool) bool {
	var direction string
	for i := 1; i < len(levels); i++ {
		prev := levels[i-1]
		element := levels[i]
		diff := element - prev
		absDiff := math.Abs(float64(diff))
		if direction == "" {
			if prev < element {
				direction = "up"
			} else {
				direction = "down"
			}
		}
		if (absDiff < 1 || absDiff > 3) || direction == "up" && diff < 0 || direction == "down" && diff > 0 {
			if dampen {
				withoutElement := DeleteIndex(levels, i)
				withoutPrev := DeleteIndex(levels, i-1)
				withoutFirst := DeleteIndex(levels, 0)
				return analyseLine(withoutElement, false) || analyseLine(withoutPrev, false) || analyseLine(withoutFirst, false)
			}
			return false
		}
	}
	return true
}

func (d Day2) Run(filepath string) {
	ans := 0
	StreamFileInts(filepath, func(levels []int) {
		val := analyseLine(levels, true)
		if val {
			ans++
		}
	})
	fmt.Println(ans)
}
