package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Runner interface {
	Run(filepath string)
}

func StreamFile(filePath string, callback func(string)) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		callback(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}

func StreamFileColumns(filepath string, callback func(string, string)) {
	StreamFile(filepath, func(line string) {
		spaceSeparated := strings.Split(line, " ")
		withoutBlank := []string{}
		for _, s := range spaceSeparated {
			if s != "" {
				withoutBlank = append(withoutBlank, s)
			}
		}
		if len(withoutBlank) != 2 {
			log.Fatalf("Expected 2 columns, got %d", len(withoutBlank))
			return
		}
		callback(withoutBlank[0], withoutBlank[1])
	})
}

func StreamFileColumnsInt(filepath string, callback func(int, int)) {
	StreamFileColumns(filepath, func(str1 string, str2 string) {
		num1 := 0
		fmt.Sscanf(str1, "%d", &num1)
		num2 := 0
		fmt.Sscanf(str2, "%d", &num2)
		callback(num1, num2)
	})
}

func GetFileName(dayNumber int, isTest bool) string {
	name := fmt.Sprintf("data/day%d", dayNumber)
	if isTest {
		return name + ".test"
	}
	return name
}

func FrequencyMap[T comparable](list []T) map[T]int {
	freqMap := make(map[T]int)
	for _, item := range list {
		freqMap[item]++
	}
	return freqMap
}
