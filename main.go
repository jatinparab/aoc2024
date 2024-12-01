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
	runner.Run(internal.GetFileName(day, isTest))
	fmt.Println("Time taken:", time.Since(t))
}
