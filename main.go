package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "csv file in format of 'question,answer")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file %s\n", *csvFileName))
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse given CSV file")
	}
	problems := parseLines(lines)

	correct := 0
	timer := time.NewTimer(time.Duration(2 * time.Second))
loop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		anserChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			anserChan <- answer
		}()
		select {
		case <-timer.C:

			break loop
		case answer := <-anserChan:
			if answer == problem.answer {
				correct++
			}
		}
	}
	fmt.Printf("\nYou socre %d/%d\n", correct, len(problems))
}

type Problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
