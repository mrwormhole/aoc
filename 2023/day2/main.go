package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const filename = "../testdata/day2_input.txt"

// First Star = 1867
// Second Star = 84538

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

	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		gameParts := strings.Split(line, ":")
		if len(gameParts) != 2 {
			log.Fatal("gameParts is not equal to 2")
		}

		game := make(map[string]int)
		sets := strings.Split(gameParts[1], ";")
		for _, set := range sets {
			colors := strings.Split(set, ",")
			for _, c := range colors {
				c = strings.TrimSpace(c)
				num, err := strconv.Atoi(strings.Split(c, " ")[0])
				if err != nil {
					log.Fatalf("strconv.Atoi(%v): %v", strings.Split(c, " ")[0], err)
				}

				var color string
				switch {
				case strings.HasSuffix(c, "red"):
					color = "red"
				case strings.HasSuffix(c, "green"):
					color = "green"
				case strings.HasSuffix(c, "blue"):
					color = "blue"
				}
				previous := game[color]
				if num > previous {
					game[color] = num
				}
			}
		}

		const (
			desiredRed   = 12
			desiredGreen = 13
			desiredBlue  = 14
		)
		if game["red"] > desiredRed || game["green"] > desiredGreen || game["blue"] > desiredBlue {
			continue
		}
		sum += i + 1
	}
	return sum
}

func part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		gameParts := strings.Split(line, ":")
		if len(gameParts) != 2 {
			log.Fatal("gameParts is not equal to 2")
		}

		game := make(map[string]int)
		sets := strings.Split(gameParts[1], ";")
		for _, set := range sets {
			colors := strings.Split(set, ",")
			for _, c := range colors {
				c = strings.TrimSpace(c)
				num, err := strconv.Atoi(strings.Split(c, " ")[0])
				if err != nil {
					log.Fatalf("strconv.Atoi(%v): %v", strings.Split(c, " ")[0], err)
				}

				var color string
				switch {
				case strings.HasSuffix(c, "red"):
					color = "red"
				case strings.HasSuffix(c, "green"):
					color = "green"
				case strings.HasSuffix(c, "blue"):
					color = "blue"
				}
				previous := game[color]
				if num > previous {
					game[color] = num
				}
			}
		}

		sum += game["red"] * game["green"] * game["blue"]
	}
	return sum
}
