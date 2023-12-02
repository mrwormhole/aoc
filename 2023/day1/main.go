package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const filename = "../testdata/day1_input.txt"

// First Star = 54634
// Second Star = 53855

func main() {
	raw, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("os.ReadFile(%v): %v", filename, err)
	}

	lines := strings.Split(string(raw), "\n")
	fmt.Println("first sum:", part1(lines))
	fmt.Println("second sum:", part2(lines))
}

func part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var calibrationValues []string
		for _, r := range line {
			if unicode.IsDigit(r) {
				calibrationValues = append(calibrationValues, string(r))
			}
		}

		var s string
		if len(calibrationValues) > 2 {
			s = string(calibrationValues[0]) + string(calibrationValues[len(calibrationValues)-1])
		} else if len(calibrationValues) == 1 {
			s = string(calibrationValues[0]) + string(calibrationValues[0])
		} else {
			s = string(calibrationValues[0]) + string(calibrationValues[1])
		}

		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("strconv.Atoi(%v): %v", s, err)
		}

		sum += val
	}
	return sum
}

func part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var calibrationValues []string
		lineLen := len(line)
		for idx, r := range line {
			if unicode.IsDigit(r) {
				calibrationValues = append(calibrationValues, string(r))
				continue
			}

			if (lineLen-1)-idx < 2 { // requires at least 3 chars
				continue
			}

			part := line[idx:min(idx+6, lineLen)]
			switch {
			case strings.HasPrefix(part, "one"):
				calibrationValues = append(calibrationValues, "1")
			case strings.HasPrefix(part, "two"):
				calibrationValues = append(calibrationValues, "2")
			case strings.HasPrefix(part, "three"):
				calibrationValues = append(calibrationValues, "3")
			case strings.HasPrefix(part, "four"):
				calibrationValues = append(calibrationValues, "4")
			case strings.HasPrefix(part, "five"):
				calibrationValues = append(calibrationValues, "5")
			case strings.HasPrefix(part, "six"):
				calibrationValues = append(calibrationValues, "6")
			case strings.HasPrefix(part, "seven"):
				calibrationValues = append(calibrationValues, "7")
			case strings.HasPrefix(part, "eight"):
				calibrationValues = append(calibrationValues, "8")
			case strings.HasPrefix(part, "nine"):
				calibrationValues = append(calibrationValues, "9")
			}
		}

		var s string
		if len(calibrationValues) > 2 {
			s = string(calibrationValues[0]) + string(calibrationValues[len(calibrationValues)-1])
		} else if len(calibrationValues) == 1 {
			s = string(calibrationValues[0]) + string(calibrationValues[0])
		} else {
			s = string(calibrationValues[0]) + string(calibrationValues[1])
		}

		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("strconv.Atoi(%v): %v", s, err)
		}

		sum += val
	}
	return sum
}
