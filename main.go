package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jatinparab/aoc2024/internal"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <day> [--test]")
		return
	}

	cmd := os.Args[1]
	isTest := len(os.Args) > 2 && os.Args[2] == "--test"

	solutions := map[string]internal.Runner{
		"1": internal.Day1{},
		"2": internal.Day2{},
		"3": internal.NewDay3(),
		"4": internal.Day4{},
		"5": internal.Day5{},
	}
	runner, ok := solutions[cmd]
	if !ok {
		log.Fatalf("Unknown command: %s", cmd)
		fmt.Println("Usage: go run main.go <day> [--test]")
		return
	}

	var day int
	fmt.Sscanf(cmd, "%d", &day)
	t := time.Now()
	fmt.Printf("-- Running solution for day %d --\n", day)
	runner.Run(internal.GetFileName(day, isTest))
	fmt.Println("Time taken:", time.Since(t))
}
