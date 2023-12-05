package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	filename  = "../testdata/day3_input.txt"
	dotRegexp = `\.{1,}`
)

var dotPattern = regexp.MustCompile(dotRegexp)

// First Star = 533784
// Second Star = 78826761

func main() {
	raw, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("os.ReadFile(%v): %v", filename, err)
	}

	lines := strings.Split(string(raw), "\n")
	fmt.Println("first sum:", part1(lines))
	fmt.Println("second sum:", part2(lines))
}

func isSymbol(b byte) bool {
	return b != '.' && !unicode.IsNumber(rune(b))
}

func part1(lines []string) int {
	sum := 0

	for lineIdx, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		lastIdx := len(line) - 1
		numberHolder := ""
		var numberPlacementIdxes []int

		for idx, c := range line {
			if unicode.IsNumber(c) {
				// add the first additional number placement padding
				if numberHolder == "" && idx != 0 {
					numberPlacementIdxes = append(numberPlacementIdxes, idx-1)
				}

				numberHolder += string(c)
				numberPlacementIdxes = append(numberPlacementIdxes, idx)

				if idx != lastIdx {
					continue
				}
			}

			if numberHolder != "" && (!unicode.IsNumber(c) || idx == lastIdx) {
				// add the last additional number placement padding
				if idx != len(line)-1 {
					numberPlacementIdxes = append(numberPlacementIdxes, idx)
				}

				var isPart bool

				// horizontal check
				start, end := numberPlacementIdxes[0], numberPlacementIdxes[len(numberPlacementIdxes)-1]
				if isSymbol(line[start]) {
					isPart = true
				}
				if isSymbol(line[end]) {
					isPart = true
				}

				// vertical check
				for _, numberPlacementIdx := range numberPlacementIdxes {
					if isPart {
						break
					}

					if lineIdx == 0 {
						isPart = isSymbol(lines[lineIdx+1][numberPlacementIdx])
						continue
					}
					if lineIdx == len(lines)-1 {
						isPart = isSymbol(lines[lineIdx-1][numberPlacementIdx])
						continue
					}

					isPart = isSymbol(lines[lineIdx-1][numberPlacementIdx]) || isSymbol(lines[lineIdx+1][numberPlacementIdx])
				}

				if !isPart {
					numberHolder = ""
					numberPlacementIdxes = nil
					continue
				}

				n, err := strconv.Atoi(numberHolder)
				if err != nil {
					log.Fatalf("strconv.Atoi(%v): %v\n", n, err)
				}

				sum += n
				numberHolder = ""
				numberPlacementIdxes = nil
			}
		}
	}
	return sum
}

func detect(line string, midTopIdx int) []int {
	lastIdx := len(line) - 1
	numberHolder := ""
	var numbers []int
	startIndex, endIndex := -1, -1

	for i, c := range line {
		if unicode.IsNumber(c) {
			if numberHolder == "" {
				startIndex = i
			}

			numberHolder += string(c)
			if i != lastIdx {
				continue
			}
		}

		if numberHolder != "" && (!unicode.IsNumber(c) || i == lastIdx) {
			if i == lastIdx {
				endIndex = i
			} else {
				endIndex = i - 1
			}

			n, err := strconv.Atoi(numberHolder)
			if err != nil {
				log.Fatalf("strconv.Atoi(%v): %v\n", n, err)
			}

			evaluateEnd := endIndex-midTopIdx == -1 || endIndex-midTopIdx == 0 || endIndex-midTopIdx == 1
			evaluateStart := startIndex-midTopIdx == -1 || startIndex-midTopIdx == 0 || startIndex-midTopIdx == 1
			if evaluateEnd || evaluateStart {
				numbers = append(numbers, n)
			}

			numberHolder = ""
			startIndex, endIndex = -1, -1
		}
	}

	return numbers
}

func part2(lines []string) int {
	sum := 0

	for lineIdx, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		for idx, c := range line {
			if c != '*' {
				continue
			}

			var vals []int

			if idx > 0 && idx < len(line)-1 { // horizontal check
				leftElems, rightElems := dotPattern.Split(line[:idx], -1), dotPattern.Split(line[idx+1:], -1)

				leftElem := leftElems[len(leftElems)-1]
				if leftElem != "" && unicode.IsNumber(rune(leftElem[len(leftElem)-1])) {
					leftNumber, err := strconv.Atoi(leftElems[len(leftElems)-1])
					if err != nil {
						log.Fatalf("strconv.Atoi(left = %v): %v\n", leftNumber, err)
					}
					vals = append(vals, leftNumber)
				}

				rightElem := rightElems[0]
				if rightElem != "" && unicode.IsNumber(rune(rightElem[0])) {
					rightNumber, err := strconv.Atoi(rightElems[0])
					if err != nil {
						log.Fatalf("strconv.Atoi(right = %v): %v\n", rightNumber, err)
					}
					vals = append(vals, rightNumber)
				}
			}

			if lineIdx > 0 && lineIdx < len(lines) { // vertical check
				topExists := unicode.IsNumber(rune(lines[lineIdx-1][idx-1])) || unicode.IsNumber(rune(lines[lineIdx-1][idx])) || unicode.IsNumber(rune(lines[lineIdx-1][idx+1]))
				botExists := unicode.IsNumber(rune(lines[lineIdx+1][idx-1])) || unicode.IsNumber(rune(lines[lineIdx+1][idx])) || unicode.IsNumber(rune(lines[lineIdx+1][idx+1]))
				if topExists {
					vals = append(vals, detect(lines[lineIdx-1], idx)...)
				}

				if botExists {
					vals = append(vals, detect(lines[lineIdx+1], idx)...)
				}
			}

			if len(vals) == 2 {
				m := 1
				for _, v := range vals {
					m = m * v
				}
				sum += m
			}
		}
	}
	return sum
}
