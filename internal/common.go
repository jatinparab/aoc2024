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

func DeleteIndex(slice []int, index int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	newSlice = append(newSlice[:index], newSlice[index+1:]...)
	return newSlice
}

func StreamFileInts(filepath string, callback func([]int)) {
	StreamFile(filepath, func(line string) {
		nums := strings.Split(line, " ")
		ints := []int{}
		for _, num := range nums {
			if num == "" {
				continue
			}
			numInt := 0
			fmt.Sscanf(num, "%d", &numInt)
			ints = append(ints, numInt)
		}
		callback(ints)
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

func FindAllIndexes(str string, substr string) []int {
	indexes := []int{}
	for i := 0; i < len(str); i++ {
		if strings.HasPrefix(str[i:], substr) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// Combinations generates all combinations of size n from the input slice
func Combinations[T any](input []T, n int) [][]T {
	if n > len(input) || n <= 0 {
		return nil
	}

	var result [][]T
	var comb func(start int, temp []T)

	comb = func(start int, temp []T) {
		if len(temp) == n {
			// Create a copy of temp to avoid mutation
			combination := make([]T, len(temp))
			copy(combination, temp)
			result = append(result, combination)
			return
		}

		for i := start; i < len(input); i++ {
			comb(i+1, append(temp, input[i]))
		}
	}

	comb(0, []T{})
	return result
}
