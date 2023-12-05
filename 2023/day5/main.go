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

const filename = "../testdata/day5_input.txt"

// First Star = 111627841
// Second Star = 69323688

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
	var seeds []int
	var attribNames []string
	type Attrib struct {
		src, dest, rangeLength int
	}

	attribs := make(map[string][]Attrib)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			line = strings.TrimSpace(strings.TrimPrefix(line, "seeds:"))
			for _, n := range strings.Split(line, " ") {
				num, err := strconv.Atoi(n)
				if err != nil {
					log.Fatalf("strconv.Atoi(%v): %v", n, err)
				}
				seeds = append(seeds, num)
			}
			continue
		}

		if strings.Contains(line, "map") {
			attribNames = append(attribNames, line)
			continue
		}

		ranges := strings.Split(line, " ")
		if len(ranges) != 3 {
			log.Fatal("len(ranges) has to be equal to 3")
		}
		rangeLength, err := strconv.Atoi(ranges[2])
		if err != nil {
			log.Fatalf("strconv.Atoi(ranges[2] = %v): %v", ranges[2], err)
		}
		src, err := strconv.Atoi(ranges[1])
		if err != nil {
			log.Fatalf("strconv.Atoi(ranges[1] = %v): %v", ranges[1], err)
		}
		dest, err := strconv.Atoi(ranges[0])
		if err != nil {
			log.Fatalf("strconv.Atoi(ranges[0] = %v): %v", ranges[0], err)
		}

		latestAttrib := attribNames[len(attribNames)-1]

		attribs[latestAttrib] = append(attribs[latestAttrib], Attrib{
			src:         src,
			dest:        dest,
			rangeLength: rangeLength,
		})
	}

	min := math.Inf(1)
	for _, s := range seeds {
		temp := s

		for _, n := range attribNames {
			attribs, ok := attribs[n]
			if !ok {
				log.Fatalf("attribs[%v] doesn't exist\n", n)
			}

			attribIdx := -1
			for i, attrib := range attribs {
				if temp >= attrib.src && temp <= attrib.src+attrib.rangeLength {
					attribIdx = i
				}
			}
			if attribIdx != -1 {
				temp = attribs[attribIdx].dest + temp - attribs[attribIdx].src
			}
		}

		if float64(temp) < min {
			min = float64(temp)
		}
	}

	return int(min)
}

// TODO: optimize these all maps into 1 map so that for each big seed, we can create optimal locations list and pick the lowest
func part2(lines []string) int {
	// type BigSeed struct {
	// 	start, end int
	// }

	var seeds []int
	var attribNames []string
	type Attrib struct {
		src, dest, rangeLength int
	}

	attribs := make(map[string][]Attrib)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			line = strings.TrimSpace(strings.TrimPrefix(line, "seeds:"))
			nums := strings.Split(line, " ")
			for i := 0; i < len(nums); i += 2 {
				src, err := strconv.Atoi(nums[i])
				if err != nil {
					log.Fatalf("strconv.Atoi(%v): %v", src, err)
				}
				rangeLength, err := strconv.Atoi(nums[i+1])
				if err != nil {
					log.Fatalf("strconv.Atoi(%v): %v", rangeLength, err)
				}
				for s := src; s < src+rangeLength; s++ {
					seeds = append(seeds, s)
				}

				// seeds = append(seeds, BigSeed{ // inclusive range for the seeds
				// 	start: src,
				// 	end:   src + rangeLength - 1,
				// })
			}
			continue
		}

		if strings.Contains(line, "map") {
			attribNames = append(attribNames, line)
			continue
		}

		ranges := strings.Split(line, " ")
		if len(ranges) != 3 {
			log.Fatal("len(ranges) has to be equal to 3")
		}
		rangeLength, err := strconv.Atoi(ranges[2])
		if err != nil {
			log.Fatalf("strconv.Atoi(ranges[2] = %v): %v", ranges[2], err)
		}
		src, err := strconv.Atoi(ranges[1])
		if err != nil {
			log.Fatalf("strconv.Atoi(ranges[1] = %v): %v", ranges[1], err)
		}
		dest, err := strconv.Atoi(ranges[0])
		if err != nil {
			log.Fatalf("strconv.Atoi(ranges[0] = %v): %v", ranges[0], err)
		}

		latestAttrib := attribNames[len(attribNames)-1]

		attribs[latestAttrib] = append(attribs[latestAttrib], Attrib{
			src:         src,
			dest:        dest,
			rangeLength: rangeLength,
		})
	}

	seeds = slices.Compact(seeds)

	min := math.Inf(1)
	for _, s := range seeds {
		temp := s

		for _, n := range attribNames {
			attribs, ok := attribs[n]
			if !ok {
				log.Fatalf("attribs[%v] doesn't exist\n", n)
			}

			attribIdx := -1
			for i, attrib := range attribs {
				if (temp >= attrib.src) && (temp <= attrib.src+attrib.rangeLength-1) {
					attribIdx = i
				}
			}

			if attribIdx != -1 {
				temp = attribs[attribIdx].dest + temp - attribs[attribIdx].src
			}
		}

		if float64(temp) < min {
			min = float64(temp)
		}
	}

	return int(min)
}
