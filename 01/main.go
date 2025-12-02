package main

import (
	"flag"
	"fmt"
	"os"
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

func part1(input string) {
	lines := ReadLines(input)
	ans := 0
	dial := 50

	for _, line := range lines {
		direction := line[0]
		value, _ := strconv.Atoi(line[1:])
		value = value % 100
		switch direction {
		case 'L':
			dial -= value
		case 'R':
			dial += value
		}

		// wrap around
		dial = (dial + 100) % 100

		// count how many times we end on 0
		if dial == 0 {
			ans++
		}
	}

	fmt.Println("Answer is", ans)
}

func part2(input string) {
	lines := ReadLines(input)
	ans := 0
	dial := 50

	for _, line := range lines {
		direction := line[0]
		value, _ := strconv.Atoi(line[1:])
		loops := value / 100
		value = value % 100
		switch direction {
		case 'L':
			dial -= value
		case 'R':
			dial += value
		}

		// count number of times we pass 0
		ans += loops
		if dial < 0 || dial >= 100 {
			ans++
		}

		// wrap around
		dial = (dial + 100) % 100
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
