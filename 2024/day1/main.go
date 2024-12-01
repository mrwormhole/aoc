package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const filename = "./input.txt"

// First Star = 1941353
// Second Star = 22539317

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
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, l := range lines {
		elems := strings.Split(l, "   ")

		n1, err := strconv.Atoi(elems[0])
		if err != nil {
			log.Fatalf("strconv.Atoi(%v): %v", elems[0], err)
		}
		left[i] = n1

		n2, err := strconv.Atoi(elems[1])
		if err != nil {
			log.Fatalf("strconv.Atoi(%v): %v", elems[1], err)
		}
		right[i] = n2
	}

	sum := 0
	for i := len(lines); i > 0; i-- {
		a, b := slices.Min(left), slices.Min(right)

		if leftIndex := slices.Index(left, a); leftIndex != -1 {
			left = slices.Delete(left, leftIndex, leftIndex+1)
		}

		if rightIndex := slices.Index(right, b); rightIndex != -1 {
			right = slices.Delete(right, rightIndex, rightIndex+1)
		}

		sum += int(math.Abs(float64(b - a)))
	}
	return sum
}

func part2(lines []string) int {
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, l := range lines {
		elems := strings.Split(l, "   ")

		n1, err := strconv.Atoi(elems[0])
		if err != nil {
			log.Fatalf("strconv.Atoi(%v): %v", elems[0], err)
		}
		left[i] = n1

		n2, err := strconv.Atoi(elems[1])
		if err != nil {
			log.Fatalf("strconv.Atoi(%v): %v", elems[1], err)
		}
		right[i] = n2
	}

	occurences := make(map[int]int)
	for _, l := range left {
		for _, r := range right {
			if l == r {
				occurence, ok := occurences[l]
				if !ok {
					occurences[l] = 1
					continue
				}
				occurences[l] = occurence + 1
			}
		}
	}

	sum := 0
	for k, v := range occurences {
		sum += k * v
	}
	return sum
}
