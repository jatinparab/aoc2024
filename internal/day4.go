package internal

import (
	"fmt"
	"strings"
)

type Day4 struct{}

func (d Day4) Run(filepath string) {
	words := [][]string{}
	StreamFile(filepath, func(line string) {
		words = append(words, strings.Split(line, ""))
	})
	fmt.Println(words)
}
