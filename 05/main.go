package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ReadLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

type Range struct {
	from int
	to   int
}

func ParseLines(lines []string) ([]Range, []int) {
	var ranges []Range
	var ids []int
	p := false

	for _, line := range lines {
		if line == "" {
			p = true
			continue
		}

		if !p {
			// e.g. 3-5
			splits := strings.Split(line, "-")
			from, _ := strconv.Atoi(splits[0])
			to, _ := strconv.Atoi(splits[1])
			ranges = append(ranges, Range{from: from, to: to})
		} else {
			// e.g. 7
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}

	return ranges, ids
}

func part1(input string) {
	lines := ReadLines(input)
	ranges, ids := ParseLines(lines)

	ans := 0
	for _, id := range ids {
		fresh := false
		for _, r := range ranges {
			if id >= r.from && id <= r.to {
				fresh = true
				break
			}
		}
		if fresh {
			ans++
		}
	}

	fmt.Println("Answer is", ans)
}

type Bound struct {
	value int
	isLow bool
}

func part2(input string) {
	lines := ReadLines(input)
	ranges, _ := ParseLines(lines)

	var bounds []Bound
	for _, r := range ranges {
		bounds = append(bounds, Bound{value: r.from, isLow: true})
		bounds = append(bounds, Bound{value: r.to, isLow: false})
	}

	// Sort bounds
	sort.Slice(bounds, func(i, j int) bool {
		if bounds[i].value == bounds[j].value {
			return bounds[i].isLow
		}
		return bounds[i].value < bounds[j].value
	})

	ans := 0
	var stack []int
	for _, b := range bounds {
		if b.isLow {
			stack = append(stack, b.value)
		} else {
			if len(stack) == 1 {
				ans += b.value - stack[0] + 1
			}
			stack = stack[:len(stack)-1]
		}
	}

	fmt.Println("Answer is", ans)
}

func main() {
	isSample := flag.Bool("s", false, "use sample input")
	flag.Parse()

	input := "input.txt"
	if *isSample {
		input = "input_sample.txt"
	}

	var start time.Time
	start = time.Now()
	part1(input)
	fmt.Println("Part 1 finished in:", time.Since(start))

	start = time.Now()
	part2(input)
	fmt.Println("Part 2 finished in:", time.Since(start))
}
