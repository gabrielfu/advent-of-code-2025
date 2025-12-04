package main

import (
	"flag"
	"fmt"
	"os"
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

func ParseLine(line string) []int {
	nums := make([]int, len(line))
	for i, ch := range line {
		nums[i] = int(ch - '0')
	}
	return nums
}

func MaxJoltageOfBank(batteries []int, numDigits int) int {
	length := len(batteries)
	digits := make([]int, numDigits)
	for i := 0; i < numDigits; i++ {
		digits[i] = batteries[length-numDigits+i]
	}

	for i := len(batteries) - numDigits - 1; i >= 0; i-- {
		cand := batteries[i]

		// insert cand into sorted array
		for k := 0; k < numDigits; k++ {
			if cand >= digits[k] {
				digits[k], cand = cand, digits[k]
			} else {
				break
			}
		}
	}

	// convert digits to number
	result := 0
	for i := 0; i < numDigits; i++ {
		result = result*10 + digits[i]
	}
	return result
}

func part1(input string) {
	lines := ReadLines(input)
	ans := 0

	for _, line := range lines {
		batteries := ParseLine(line)
		joltage := MaxJoltageOfBank(batteries, 2)
		ans += joltage
	}
	fmt.Println("Answer is", ans)
}

func part2(input string) {
	lines := ReadLines(input)
	ans := 0

	for _, line := range lines {
		batteries := ParseLine(line)
		joltage := MaxJoltageOfBank(batteries, 12)
		ans += joltage
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
