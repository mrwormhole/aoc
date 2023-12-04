package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const filename = "../testdata/day4_input.txt"

// First Star = 27059
// Second Star = 5744979

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

		line = strings.TrimSpace(line[strings.Index(line, ":")+1:])
		parts := strings.Split(line, "|")
		winningNumbers, numbers := strings.Split(parts[0], " "), strings.Split(parts[1], " ")

		winners := make(map[int]struct{})
		for _, n := range winningNumbers {
			if strings.TrimSpace(n) == "" {
				continue
			}

			num, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("strconv.Atoi(%v): %v", n, err)
			}

			winners[num] = struct{}{}
		}

		match := 0
		for _, n := range numbers {
			if strings.TrimSpace(n) == "" {
				continue
			}

			num, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("strconv.Atoi(%v): %v", n, err)
			}

			if _, ok := winners[num]; ok {
				match++
			}
		}

		if match > 0 {
			sum += int(math.Pow(2, float64(match-1)))
		}
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	scratchCardCounts := make(map[string]int)

	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		line = strings.TrimSpace(line[strings.Index(line, ":")+1:])
		parts := strings.Split(line, "|")
		winningNumbers, numbers := strings.Split(parts[0], " "), strings.Split(parts[1], " ")

		winners := make(map[int]struct{})
		for _, n := range winningNumbers {
			if strings.TrimSpace(n) == "" {
				continue
			}

			num, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("strconv.Atoi(%v): %v", n, err)
			}

			winners[num] = struct{}{}
		}

		match := 0
		for _, n := range numbers {
			if strings.TrimSpace(n) == "" {
				continue
			}

			num, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("strconv.Atoi(%v): %v", n, err)
			}

			if _, ok := winners[num]; ok {
				match++
			}
		}

		copyCount := scratchCardCounts["cardIdx "+strconv.Itoa(i)]
		scratchCardCounts["cardIdx "+strconv.Itoa(i)] = copyCount + 1
		sum += copyCount + 1

		for m := 1; m < match+1; m++ {
			scratchCardCounts["cardIdx "+strconv.Itoa(i+m)] += 1
			if copyCount > 0 {
				scratchCardCounts["cardIdx "+strconv.Itoa(i+m)] += copyCount
			}
		}
	}
	return sum
}
